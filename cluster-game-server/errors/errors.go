package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthAPIThrowError(v string) error {
	return status.Errorf(codes.Unauthenticated, "Auth API Rejected Session Token: %s", v)
}

func AuthMDNotFound() error {
	return status.Errorf(codes.Unauthenticated, "session token not found in header")
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
