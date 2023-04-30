package me

import v1 "github.com/kyfk/go-grpc-sqlc-boilerplate/protogen/proto/me/v1"

type Server struct {
	v1.UnimplementedMeServer
}

func New() (*Server, error) {
	return &Server{}, nil
}
