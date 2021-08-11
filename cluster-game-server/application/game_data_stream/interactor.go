package application

import (
	"io"
	"sync"
	"time"

	account "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	character "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/character"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"google.golang.org/grpc/metadata"
)

type interactor struct {
	accountrepo   account.Repository
	characterrepo character.Repository
	streams       *[]pb.GameServices_GameDataStreamServer
	smu           sync.Mutex
}

func New(ar account.Repository, cr character.Repository) InputPort {
	i := &interactor{
		accountrepo:   ar,
		characterrepo: cr,
		streams:       &[]pb.GameServices_GameDataStreamServer{},
		smu:           sync.Mutex{},
	}
	return i
}

func (i *interactor) Handle(input InputData) OutputData {
	stream := input.Stream

	//ヘッダーチェック
	headers, ok := metadata.FromIncomingContext(stream.Context())
	if !ok && len(headers["sessiontoken"]) > 0 {
		return OutputData{
			Err: errors.AuthMDNotFound(),
		}
	}

	token := headers["sessiontoken"][0]

	//tokenチェック
	val, err := i.accountrepo.Find(token)
	if err != nil {
		return OutputData{Err: err}
	}

	userid := val.ID

	if status := i.accountrepo.GetSessionStatus(userid); !status.IsConnectable() {
		return OutputData{Err: errors.InvalidOperation()}
	}

	//接続処理
	i.accountrepo.Connect(userid)

	errch := make(chan error)

	//送信用ストリームリストに登録
	i.smu.Lock()
	if len(*i.streams) == 0 {
		go i.sender()
	}
	*i.streams = append(*i.streams, stream)
	i.smu.Unlock()

	defer func() {
		//ストリームリストから削除
		i.smu.Lock()

		var res []pb.GameServices_GameDataStreamServer
		for _, v := range *i.streams {
			if v != stream {
				res = append(res, v)
			}
		}
		*i.streams = res

		i.smu.Unlock()
	}()

	//受信開始
	go i.receiver(stream, errch, userid)

	err = <-errch
	if err == io.EOF {
		i.accountrepo.Disconnect(userid)
		return OutputData{}
	}

	//TODO: エラー用のステータスに変更
	i.accountrepo.Disconnect(userid)
	return OutputData{
		Err: err,
	}
}

func (i *interactor) receiver(stream pb.GameServices_GameDataStreamServer, errch chan<- error, id account.ID) {

	for {
		req, err := stream.Recv()
		if err != nil {
			errch <- err
			break
		}

		if status := i.accountrepo.GetSessionStatus(id); !status.IsActive() {
			errch <- errors.SessionNotActive()
			break
		}

		message := req.Message

		switch x := message.(type) {
		case *pb.GameDataRequest_CharacterUpdate:
			c := x.CharacterUpdate
			i.characterrepo.Update(c)
		}
	}
}

func (i *interactor) sender() {
	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()
	for {
		<-t.C

		clist := i.characterrepo.GetAll()
		if clist == nil {
			continue
		}
		characters := pb.Characters{
			Characters: *clist,
		}
		response := pb.GameDataResponse{
			Message: &pb.GameDataResponse_CharacterDatas{
				CharacterDatas: &characters,
			},
		}
		i.smu.Lock()
		if len(*i.streams) == 0 {
			break
		}
		for _, s := range *i.streams {
			if err := s.Send(&response); err != nil {
				continue
			}
		}
		i.smu.Unlock()
	}
}
