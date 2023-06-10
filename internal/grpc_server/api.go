package grpc_server

import (
	"github.com/lowl11/lazylog/log"
	"google.golang.org/grpc"
	"net"
)

func (server *Server) Start(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err, "Create TCP listener for gRPC server error")
	}
	server.listener = listener

	log.Info("gRPC server started at", port)
	if err = server.server.Serve(listener); err != nil {
		log.Fatal(err, "Run gRPC server error")
	}
}

func (server *Server) Close() error {
	if server.listener == nil {
		return nil
	}

	return server.listener.Close()
}

func (server *Server) Get() *grpc.Server {
	return server.server
}
