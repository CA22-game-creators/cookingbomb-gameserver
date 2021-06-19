package main

import (
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/api"
)

type game struct{}

func main() {
	// listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))

	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// log.Printf("Listening :%s", os.Getenv("PORT"))

	// grpcServer := grpc.NewServer()
	// service := &controller.GameService{}
	// pb.RegisterGameServicesServer(grpcServer, service)
	// reflection.Register(grpcServer)

	// if err = grpcServer.Serve(listen); err != nil {
	// 	log.Fatalf(err.Error())
	// }

	api.GetAccount("hoghoeg")
}
