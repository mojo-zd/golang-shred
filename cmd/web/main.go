package main

import (
	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/sirupsen/logrus"
	"net/http"
)
func main() {
	app := iris.New()
	app.Handle(http.MethodGet, "/charts", func(ctx context.Context) {
		_, err := ctx.JSON(map[string]interface{}{"name":"value"})
		if err != nil {
			logrus.Error(err)
		}
	})
	err := app.Run(iris.Addr(":8080"))
	logrus.Info(err)
}
