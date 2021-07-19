package testdata

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"google.golang.org/grpc"
)

type DummyStream struct {
	grpc.ServerStream
}

func (x *DummyStream) Send(m *pb.GameDataResponse) error {
	return nil
}

func (x *DummyStream) Recv() (*pb.GameDataRequest, error) {
	return nil, nil
}
