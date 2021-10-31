package testdata

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"google.golang.org/grpc"
)

type Stream struct {
	grpc.ServerStream
}

func (x *Stream) Send(m *pb.GameDataResponse) error {
	return nil
}

func (x *Stream) Recv() (*pb.GameDataRequest, error) {
	return nil, nil
}
