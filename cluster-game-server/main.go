package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Listening :%s", os.Getenv("PORT"))
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf(err.Error())
	}
}
