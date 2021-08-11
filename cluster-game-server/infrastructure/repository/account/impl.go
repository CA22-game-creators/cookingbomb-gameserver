package infra

import (
	"context"
	"log"
	"time"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
	"github.com/oklog/ulid/v2"
	goCache "github.com/patrickmn/go-cache"

	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/api"
)

type impl struct {
	instance *goCache.Cache
}

func New(c *goCache.Cache) domain.Repository {
	return &impl{
		instance: c,
	}
}

func (i impl) Find(sesisonToken string) (domain.Account, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*5,
	)
	defer cancel()
	conn, err := api.New()
	if err != nil {
		log.Print("API Server Connection Failed: ", err)
		return domain.Account{}, errors.APIConnectionLost()
	}
	defer conn.Close()

	client := pb.NewAccountServicesClient(conn)
	req := &pb.GetAccountInfoRequest{SessionToken: sesisonToken}
	res, err := client.GetAccountInfo(ctx, req)
	if err != nil {
		return domain.Account{}, errors.AuthAPIThrowError(err.Error())
	}

	return domain.FromRepository(
		res.AccountInfo.Id,
		res.AccountInfo.Name,
	), nil
}

func (i impl) GetSessionStatus(id domain.ID) domain.StatusEnum {
	status, ok := i.instance.Get(ulid.ULID(id).String())
	if !ok {
		return domain.UNSPECIFIED
	}
	return status.(domain.StatusEnum)
}

func (i impl) Connect(id domain.ID) {
	i.instance.Set(ulid.ULID(id).String(), domain.CONNECTED, goCache.NoExpiration)
}

func (i impl) Disconnect(id domain.ID) {
	i.instance.Set(ulid.ULID(id).String(), domain.DISCONNECTED_BY_CLIENT, goCache.DefaultExpiration)
}
