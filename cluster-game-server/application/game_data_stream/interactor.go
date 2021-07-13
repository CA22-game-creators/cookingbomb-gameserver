package application

import (
	"io"
	"sync"
	"time"

	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// sessionToken: characterStatus
var handler = struct {
	characters map[string]*pb.Character
	mu         sync.Mutex
}{}

type interactor struct {
	repository domain.Repository
}

func New(r domain.Repository) InputPort {
	return &interactor{
		repository: r,
	}
}

func (i interactor) Handle(input InputData) OutputData {

	var errch chan error

	go i.receiver(input.Stream, errch)
	go i.sender(input.Stream, errch)

	return OutputData{Err: <-errch}
}

func (i interactor) receiver(stream pb.GameServices_GameDataStreamServer, errch chan error) {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			continue
		}
		if err != nil {
			errch <- err
		}

		if status := i.repository.GetSessionStatus(req.GetSessionToken()); !status.IsActive() {
			errch <- errors.SessionNotActive()
		}

		handler.mu.Lock()
		handler.characters[req.SessionToken] = req.GetCharacterUpdate()
		handler.mu.Unlock()
	}
}

func (i interactor) sender(stream pb.GameServices_GameDataStreamServer, errch chan error) {
	for {
		time.Sleep(100 * time.Millisecond) // 10fps
		if len(handler.characters) == 0 {
			continue
		}

		res := make([]*pb.Character, 0, len(handler.characters))
		for _, v := range handler.characters {
			res = append(res, v)
		}

		err := stream.Send(
			&pb.GameDataResponse{
				Message: &pb.GameDataResponse_CharacterDatas{
					CharacterDatas: &pb.Characters{
						Characters: res,
					},
				},
				ServerTime: timestamppb.Now(),
			},
		)
		if err != nil {
			errch <- err
		}
	}
}
