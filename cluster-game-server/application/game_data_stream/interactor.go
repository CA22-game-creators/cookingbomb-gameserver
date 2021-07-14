package application

import (
	"io"
	"sync"
	"time"

	account "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	character "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/character"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

type interactor struct {
	accountrepo   account.Repository
	characterrepo character.Repository
	streams       []pb.GameServices_GameDataStreamServer
	smu           *sync.Mutex
}

func New(ar account.Repository, cr character.Repository) InputPort {
	i := &interactor{
		accountrepo:   ar,
		characterrepo: cr,
		streams:       []pb.GameServices_GameDataStreamServer{},
		smu:           &sync.Mutex{},
	}
	go sender(*i)
	return i
}

func (i *interactor) Handle(input InputData) {
	stream := input.Stream

	i.smu.Lock()
	i.streams = append(i.streams, stream)
	i.smu.Unlock()

	defer func() {
		i.smu.Lock()

		var res []pb.GameServices_GameDataStreamServer
		for _, v := range i.streams {
			if v != stream {
				res = append(res, v)
			}
		}
		i.streams = res

		i.smu.Unlock()
	}()

	cindex := -1

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			continue
		}
		if err != nil {
			break
		}

		token := req.GetSessionToken()
		if status := i.accountrepo.GetSessionStatus(token); !status.IsActive() {
			break
		}

		message := req.Message

		switch x := message.(type) {
		case *pb.GameDataRequest_CharacterUpdate:
			c := x.CharacterUpdate
			if cindex == -1 {
				cindex = i.characterrepo.Add(c)
			} else {
				i.characterrepo.Update(c, cindex)
			}
		}
	}
}

func sender(i interactor) {
	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()
	for {
		<-t.C

		clist := i.characterrepo.GetAll()
		if clist == nil {
			continue
		}

		characters := pb.Characters{}
		copy(characters.Characters, *clist)
		responce := pb.GameDataResponse{
			Message: &pb.GameDataResponse_CharacterDatas{
				CharacterDatas: &characters,
			},
		}

		i.smu.Lock()
		for _, s := range i.streams {
			s.Send(&responce)
		}
		i.smu.Unlock()
	}
}
