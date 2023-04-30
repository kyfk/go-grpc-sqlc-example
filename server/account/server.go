package account

import v1 "github.com/kyfk/go-grpc-sqlc-boilerplate/protogen/proto/account/v1"

type Server struct {
	v1.UnimplementedAccountServer
}

func New() (*Server, error) {
	return &Server{}, nil
}
