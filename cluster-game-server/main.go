package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	controller "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/presentation/game"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Listening :%s", os.Getenv("PORT"))

	grpcServer := grpc.NewServer()
	service := &controller.Service{}
	pb.RegisterGameServicesServer(grpcServer, service)
	reflection.Register(grpcServer)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf(err.Error())
	}
}
