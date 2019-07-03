package server

import (
	"github.com/kataras/iris"
	"os"
)

func SetUp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	return app
}

func RunServer(app *iris.Application) {
	_ = app.Run(
		iris.Addr(os.Getenv("PORT")),
	)
}
