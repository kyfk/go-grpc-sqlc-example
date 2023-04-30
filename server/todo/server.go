package todo

import v1 "github.com/kyfk/go-grpc-sqlc-boilerplate/protogen/proto/todo/v1"

type Server struct {
	v1.UnimplementedTodoServiceServer
}

func New() (*Server, error) {
	return &Server{}, nil
}
