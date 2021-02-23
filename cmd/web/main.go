package main

import (
	"net/http"

	"sigs.k8s.io/controller-runtime/pkg/controller"

	"sigs.k8s.io/controller-runtime/pkg/manager"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/sirupsen/logrus"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	app := iris.New()
	app.Handle(http.MethodGet, "/charts", func(ctx context.Context) {
		_, err := ctx.JSON(map[string]interface{}{"name": "value"})
		if err != nil {
			logrus.Error(err)
		}
	})
	err := app.Run(iris.Addr(":8080"))
	logrus.Info(err)
}

func manager1() {
	mgr, _ := manager.New(config.GetConfigOrDie(), manager.Options{})
	controller.New("foo-controller", mgr, controller.Options{})
	mgr.GetWebhookServer()
}
