package application

import (
	"fmt"
	"io"
	"sync"
	"time"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	account "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	character "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/character"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
)

type interactor struct {
	accountRepo   account.Repository
	characterRepo character.Repository
	mu            sync.Mutex
}

func New(ar account.Repository, cr character.Repository) InputPort {
	return &interactor{
		accountRepo:   ar,
		characterRepo: cr,
		mu:            sync.Mutex{},
	}
}

func (i *interactor) Handle(input InputData) OutputData {
	stream := input.Stream
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok || len(md.Get("session-token")) == 0 {
		return OutputData{
			Err: errors.Unauthenticated("session-token is required in metadata"),
		}
	}
	_, err := i.accountRepo.Find(md.Get("session-token")[0])
	if err != nil {
		return OutputData{
			Err: errors.Unauthenticated(err.Error()),
		}
	}

	errch := make(chan error)
	go i.receiver(stream, errch)
	go i.sender(stream, errch)

	err = <-errch
	if err == io.EOF {
		return OutputData{}
	}
	return OutputData{
		Err: err,
	}
}

func (i *interactor) receiver(stream pb.GameServices_GameDataStreamServer, errch chan<- error) {
	var id string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			errch <- err
			break
		}
		character := req.Character
		if character.Id == "" {
			continue
		}
		if id != character.Id {
			if id != "" {
				errch <- errors.InvalidArgument(fmt.Sprintf("1プレイヤーで操作できるキャラidは1つ(%s)です", id))
				break
			}
			id = character.Id
		}
		i.characterRepo.Update(character)
	}

	for _, v := range i.characterRepo.GetAll() {
		if v.Id == id {
			i.characterRepo.Delete(v)
			return
		}
	}
}

func (i *interactor) sender(stream pb.GameServices_GameDataStreamServer, errch chan<- error) {
	t := time.NewTicker(time.Millisecond * 100)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			stream.Send(&pb.GameDataResponse{
				Characters: i.characterRepo.GetAll(),
				ServerTime: timestamppb.Now(),
			})
		case <-stream.Context().Done():
			return
		}
	}
}
