package api

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
)

func New() (*grpc.ClientConn, error) {
	host := os.Getenv("API_ADDRESS")

	if os.Getenv("ENV") == "local" {
		return grpc.Dial(
			host,
			grpc.WithInsecure(),
			grpc.WithBlock(),
		)
	}

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.Internal(err.Error())
	}
	return grpc.Dial(
		host,
		grpc.WithTransportCredentials(
			credentials.NewTLS(
				&tls.Config{
					RootCAs: systemRoots,
				},
			),
		),
	)
}
