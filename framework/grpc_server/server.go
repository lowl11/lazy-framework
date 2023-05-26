package grpc_server

import (
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func New() *Server {
	return &Server{
		server: grpc.NewServer(),
	}
}
