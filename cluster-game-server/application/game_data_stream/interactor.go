package application

import (
	"sync"
	"time"

	account "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	character "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/character"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

type interactor struct {
	repository account.Repository
	writewg    *sync.WaitGroup
	sendwg     *sync.WaitGroup
}

var characterMap map[string]character.Character
var arrayMutex = &sync.Mutex{}
var streamArray []*pb.GameServices_GameDataStreamServer

func New(r account.Repository) InputPort {
	i := &interactor{
		repository: r,
		writewg:    &sync.WaitGroup{},
		sendwg:     &sync.WaitGroup{},
	}
	go sender(*i)
	return i
}

func (i interactor) Handle(input InputData) {
	stream := input.Stream

	arrayMutex.Lock()
	streamArray = append(streamArray, &stream)
	arrayMutex.Unlock()

	defer func() {
		arrayMutex.Lock()

		var res []*pb.GameServices_GameDataStreamServer
		for _, v := range streamArray {
			if v != &stream {
				res = append(res, v)
			}
		}
		streamArray = res

		arrayMutex.Unlock()
	}()

	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}

		token := req.GetSessionToken()
		if status := i.repository.GetSessionStatus(token); !status.IsActive() {
			break
		}

		message := req.Message

		switch x := message.(type) {
		case *pb.GameDataRequest_CharacterUpdate:
			c := character.FromRepository(x.CharacterUpdate)
			i.sendwg.Wait()
			i.writewg.Add(1)
			characterMap[token] = c
			i.writewg.Done()
		}
	}
	return
}

func sender(i interactor) {
	t := time.NewTicker(33 * time.Millisecond)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			i.sendwg.Add(1)
			i.writewg.Wait()
			arrayMutex.Lock()

			arrayMutex.Unlock()
			i.sendwg.Done()
		}
	}
}
