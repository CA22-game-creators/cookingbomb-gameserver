package api

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	errors "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	account "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/model/account"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
	"google.golang.org/grpc"
)

func GetAccountInfo(token string) (account.Account, error) {
	// TODO: タイムアウトの検討: 5秒

	if flag.Lookup("test.v") != nil {
		if token == "invalid" {
			return account.Account{}, errors.AuthAPIThrowError()
		}
		return account.Account{
			ID:   "00000000-0000-0000-0000-000000000000",
			Name: "Test Name",
		}, nil
	}

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
		return account.Account{}, errors.APIConnectionLost()
	}
	defer conn.Close()

	client := pb.NewAccountServicesClient(conn)

	req := &pb.GetAccountInfoRequest{SessionToken: token}
	res, err := client.GetAccountInfo(ctx, req)
	if err != nil {
		log.Print("API Server Returned Error: ", err)
		return account.Account{}, errors.AuthAPIThrowError()
	}

	ac := account.NewFromAPIResponce(res.GetAccountInfo())
	return ac, nil
}
