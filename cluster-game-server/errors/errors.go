package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthAPIThrowError() error {
	return status.Error(codes.Unauthenticated, "Auth API Rejected Session Token")
}

func Unauthorized() error {
	return status.Error(codes.PermissionDenied, "Connection Not Permited")
}

func SessionNotActive() error {
	return status.Error(codes.PermissionDenied, "Session Not Active")
}

func InvalidOperation() error {
	return status.Error(codes.FailedPrecondition, "Operation is Invalid")
}

func InvalidArgument(err error) error {
	return status.Error(codes.InvalidArgument, "Invalid Argument Found: "+err.Error())
}

func APIConnectionLost() error {
	return status.Error(codes.Internal, "API Server Connection Failed")
}

func NoStatusFound() error {
	return status.Error(codes.NotFound, "No Connection Status Found")
}
