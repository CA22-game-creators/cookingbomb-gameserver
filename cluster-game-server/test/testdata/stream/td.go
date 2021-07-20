package testdata

import (
	"io"
	"time"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"google.golang.org/grpc"
)

type DummyStream struct {
	grpc.ServerStream
	SendLog    []*pb.GameDataResponse
	SendErr    error
	RecvExpect []*pb.GameDataRequest
	RecvErr    error
	RecvDelay  time.Duration
}

func (x *DummyStream) Send(m *pb.GameDataResponse) error {
	if x.SendErr != nil {
		return x.SendErr
	}
	x.SendLog = append(x.SendLog, m)
	return nil
}

func (x *DummyStream) Recv() (*pb.GameDataRequest, error) {
	if x.RecvDelay != 0 {
		time.Sleep(x.RecvDelay)
	}
	if x.RecvErr != nil {
		return nil, x.RecvErr
	}
	if len(x.RecvExpect) == 0 {
		return nil, io.EOF
	}
	data := x.RecvExpect[0]
	x.RecvExpect = x.RecvExpect[1:]
	return data, nil
}
