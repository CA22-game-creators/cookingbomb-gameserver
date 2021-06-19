package api

import (
	"context"
	"log"
	"time"

	account "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/model/account"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
	"google.golang.org/grpc"
)

func getAccountInfo(token string) (*pb.AccountInfo, error) {
	// TODO: タイムアウトの検討: 5秒

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*5,
	)
	defer cancel()

	conn, err := grpc.Dial(
		"localhost:8080",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second*5),
	)
	if err != nil {
		log.Print("API Server Connection Failed: ", err)
		return &pb.AccountInfo{}, err
	}
	defer conn.Close()

	client := pb.NewAccountServicesClient(conn)

	req := &pb.GetAccountInfoRequest{SessionToken: token}
	res, err := client.GetAccountInfo(ctx, req)
	if err != nil {
		log.Print("API Server Returned Error: ", err)
		return &pb.AccountInfo{}, err
	}

	return res.GetAccountInfo(), nil
}

func GetId(token string) (string, error) {
	res, err := getAccountInfo(token)
	if err != nil {
		return "", err
	}
	id := res.GetId()
	return id, nil
}

func GetAccount(token string) (account.Account, error) {
	res, err := getAccountInfo(token)
	if err != nil {
		return account.Account{}, err
	}
	return account.NewFromAPIResponce(res), nil
}
