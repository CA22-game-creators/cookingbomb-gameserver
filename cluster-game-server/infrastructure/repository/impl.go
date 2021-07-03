package infra

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
	goCache "github.com/patrickmn/go-cache"
	"google.golang.org/grpc"

	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/cache"
)

type impl struct {
}

func New() domain.Repository {
	return &impl{}
}

func (i impl) Find(sesisonToken string) (domain.Account, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*5,
	)
	defer cancel()
	conn, err := grpc.Dial(
		os.Getenv("API_ADDRESS"),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Print("API Server Connection Failed: ", err)
		return domain.Account{}, errors.APIConnectionLost()
	}
	defer conn.Close()

	client := pb.NewAccountServicesClient(conn)
	req := &pb.GetAccountInfoRequest{SessionToken: sesisonToken}
	res, err := client.GetAccountInfo(ctx, req)
	if err != nil || res.GetAccountInfo() == nil {
		log.Print("API Server Returned Error: ", err)
		return domain.Account{}, errors.AuthAPIThrowError()
	}

	return domain.FromRepository(
		res.AccountInfo.Id,
		res.AccountInfo.Name,
	), nil
}

func (i impl) GetStatus(sessionToken string) domain.StatusEnum {
	status, ok := cache.Instance.Get(sessionToken)
	if !ok {
		return domain.UNSPECIFIED
	}
	return status.(domain.StatusEnum)
}

func (i impl) Connect(sessionToken string) {
	cache.Instance.Set(sessionToken, domain.CONNECTED, goCache.NoExpiration)
}

func (i impl) Disconnect(sessionToken string) {
	cache.Instance.Set(sessionToken, domain.DISCONNECTED_BY_CLIENT, goCache.NoExpiration)
}
