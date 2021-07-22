package main

import (
	"context"
	"io"
	"log"
	"time"

	api "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
	game "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"

	"github.com/mattn/go-tty"
	"google.golang.org/grpc"
)

func makeAccount() string {
	log.Print("Start Signup and Get Session Token")

	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Connection Fail to API Server: %s", err)
	}

	defer conn.Close()

	client := api.NewAccountServicesClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	signup_req := api.SignupRequest{
		Name: "TestUser",
	}

	surep, err := client.Signup(ctx, &signup_req)
	if err != nil {
		log.Fatalf("Sign Up Failed: %s", err)
	}

	id := surep.AccountInfo.GetId()
	auth := surep.GetAuthToken()

	session_req := api.GetSessionTokenRequest{
		UserId:    id,
		AuthToken: auth,
	}

	sesrep, err := client.GetSessionToken(ctx, &session_req)
	if err != nil {
		log.Fatalf("Get Session Token Failed: %s", err)
	}

	log.Print("Session token acquired")
	log.Printf("Token: %s", sesrep.GetSessionToken())
	return sesrep.GetSessionToken()
}

func connect(token string) {
	log.Print("Start Connect to Game Server")

	address := "localhost:8085"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Connection Fail to API Server: %s", err)
	}

	defer conn.Close()

	client := game.NewGameServicesClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	req := game.ConnectionRequest{
		SessionToken: token,
	}

	rep, err := client.Connect(ctx, &req)
	if err != nil {
		log.Fatalf("Connect Failed. %s", err)
	}

	if rep.Status != game.ConnectionStatusEnum_CONNECTED {
		log.Fatalf("Connect Failed. Status: %s", rep.Status)
	}

	log.Printf("Connected. Status: %s", rep.Status)
}

func sender(token string, stream game.GameServices_GameDataStreamClient, waitc chan struct{}) {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer tty.Close()

	character := game.Character{
		Id:       "TestUser",
		IsActive: true,
		Position: &game.Position{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Rotation: &game.Rotation{
			X: 0,
			Y: 0,
			Z: 0,
			W: 0,
		},
		Verocity: &game.Verocity{
			X: 0,
			Y: 0,
			Z: 0,
		},
	}

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Print(err)
			break
		}

		switch string(r) {
		case "w":
			character.Position.Y++
		case "s":
			character.Position.Y--
		case "a":
			character.Position.X++
		case "d":
			character.Position.X--
		case "q":
			close(waitc)
			return
		}

		err = stream.Send(&game.GameDataRequest{
			SessionToken: token,
			Message: &game.GameDataRequest_CharacterUpdate{
				CharacterUpdate: &character,
			},
		})

		if err != nil {
			log.Fatalf("Data Send Failed: %v", err)
		}

		log.Printf("Data Send: %s", character.String())
	}
	close(waitc)
}

func receiver(stream game.GameServices_GameDataStreamClient, waitc chan struct{}) {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Data Recv Failed: %v", err)
		}
		log.Printf("%s", in.Message)
	}
	close(waitc)
}

func stream(token string) {
	log.Print("Start Streaming...")

	address := "localhost:8085"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Streaming Dial Failed. %s", err)
	}

	defer conn.Close()

	client := game.NewGameServicesClient(conn)
	stream, err := client.GameDataStream(context.Background())
	if err != nil {
		log.Fatalf("Streaming Failed. %s", err)
	}

	waitc := make(chan struct{})
	go sender(token, stream, waitc)
	go receiver(stream, waitc)
	<-waitc
}

func main() {
	token := makeAccount()
	connect(token)
	stream(token)
}
