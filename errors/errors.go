package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Unauthenticated(v string) error {
	return status.Errorf(codes.Unauthenticated, v)
}

func Internal(v string) error {
	return status.Errorf(codes.Internal, v)
}

func InvalidArgument(v string) error {
	return status.Errorf(codes.InvalidArgument, v)
}
