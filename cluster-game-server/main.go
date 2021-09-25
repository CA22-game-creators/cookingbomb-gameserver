package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	coresdk "agones.dev/agones/pkg/sdk"
	sdk "agones.dev/agones/sdks/go"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/di"
)

const aliveMin = 3

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("Listening :%s", os.Getenv("PORT"))

	if os.Getenv("ENV") != "local" {
		setupAgones()
	}

	grpcServer := grpc.NewServer()
	service := di.DI()
	pb.RegisterGameServicesServer(grpcServer, service)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf(err.Error())
	}
}

func setupAgones() {
	s, err := sdk.NewSDK()
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err := s.Ready(); err != nil {
		log.Fatalf(err.Error())
	}
	go healthCheck(s)
	shutdownAfterAllocation(s)
}

func healthCheck(s *sdk.SDK) {
	for range time.Tick(time.Second) {
		if err := s.Health(); err != nil {
			log.Fatalf(err.Error())
		}
	}
}

func shutdownAfterAllocation(s *sdk.SDK) {
	err := s.WatchGameServer(func(gs *coresdk.GameServer) {
		if gs.Status.State == "Allocated" {
			time.Sleep(time.Duration(aliveMin) * time.Minute)
			if err := s.Shutdown(); err != nil {
				log.Fatalf("Could not shutdown: %v", err)
			}
		}
	})
	if err != nil {
		log.Fatalf("Could not watch Game Server events, %v", err)
	}
}
