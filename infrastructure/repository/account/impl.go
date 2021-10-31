package infra

import (
	"context"
	"time"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
	"google.golang.org/grpc/metadata"

	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/api"
)

type impl struct{}

func New() domain.Repository {
	return impl{}
}

func (i impl) Find(sesisonToken string) (domain.Account, error) {
	conn, err := api.New()
	if err != nil {
		return domain.Account{}, errors.Internal("fail to prepare connection to api")
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*5,
	)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "session-token", sesisonToken)
	client := pb.NewAccountServicesClient(conn)
	res, err := client.GetAccountInfo(ctx, &pb.GetAccountInfoRequest{})
	if err != nil {
		return domain.Account{}, errors.InvalidArgument(err.Error())
	}

	return domain.FromRepository(
		res.Account.Id,
		res.Account.Name,
	), nil
}
