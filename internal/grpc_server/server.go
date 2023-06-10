package grpc_server

import (
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func New(options ...grpc.ServerOption) *Server {
	return &Server{
		server: grpc.NewServer(options...),
	}
}
