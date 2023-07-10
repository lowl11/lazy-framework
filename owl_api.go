package owl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazylog/log"
	"github.com/lowl11/owl/data/interfaces"
	"google.golang.org/grpc"
)

func (app *App) ShutdownAction(action func()) {
	app.shutdownService.Add(action)
}

func (app *App) Start(port string) {
	app.server.Start(port)
}

func (app *App) StartHttp2(port string) {
	app.server.StartHttp2(port)
}

func (app *App) StartGrpc(port string) {
	app.ShutdownAction(func() {
		if err := app.grpcServer.Close(); err != nil {
			log.Error(err, "Close gRPC server connection error")
			return
		}
		log.Info("gRPC server connection closed!")
	})

	app.getGrpcServer().Start(port)
}

func (app *App) Echo() *echo.Echo {
	return app.server.(interfaces.IEchoServer).Get()
}

func (app *App) Fiber() *fiber.App {
	return app.server.(interfaces.IFiberServer).Get()
}

func (app *App) Grpc() *grpc.Server {
	return app.getGrpcServer().Get()
}
