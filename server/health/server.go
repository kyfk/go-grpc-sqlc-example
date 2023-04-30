package health

import healthpb "google.golang.org/grpc/health/grpc_health_v1"

type Server struct {
	healthpb.UnimplementedHealthServer
}

func New() *Server {
	return &Server{}
}
