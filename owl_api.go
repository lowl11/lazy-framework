package owl

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazylog/log"
	"github.com/lowl11/owl/data/interfaces"
	"google.golang.org/grpc"
)

func (owl *Owl) ShutdownAction(action func()) {
	owl.shutdownService.Add(action)
}

func (owl *Owl) Start(port string) {
	owl.server.Start(port)
}

func (owl *Owl) StartHttp2(port string) {
	owl.server.StartHttp2(port)
}

func (owl *Owl) StartGrpc(port string) {
	owl.ShutdownAction(func() {
		if err := owl.grpcServer.Close(); err != nil {
			log.Error(err, "Close gRPC server connection error")
			return
		}
		log.Info("gRPC server connection closed!")
	})

	owl.getGrpcServer().Start(port)
}

func (owl *Owl) Echo() *echo.Echo {
	return owl.server.(interfaces.IEchoServer).Get()
}

func (owl *Owl) Grpc() *grpc.Server {
	return owl.getGrpcServer().Get()
}
