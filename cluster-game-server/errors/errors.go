package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthAPIThrowError() error {
	return status.Error(codes.Unauthenticated, "Auth API Rejected Session Token")
}

func SessionNotActive() error {
	return status.Error(codes.PermissionDenied, "Session Not Active")
}

func InvalidOperation() error {
	return status.Error(codes.FailedPrecondition, "Operation is Invalid")
}

func InvalidArgument(v string) error {
	return status.Errorf(codes.InvalidArgument, "Invalid Argument Found: %s", v)
}

func APIConnectionLost() error {
	return status.Error(codes.Internal, "API Server Connection Failed")
}
