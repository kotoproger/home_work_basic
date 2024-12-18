package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kotoproger/home_work_basic/configapp"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/app"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/app/handler"
)

func Start(app app.App, conf configapp.ConfigApp) {
	router := gin.Default()
	httpServer := &http.Server{ //nolint:gosec
		Addr:    "0.0.0.0:" + conf.GetString("http_port"),
		Handler: router,
	}

	userH := handler.NewUser(app)

	router.Use(userH.Auth)
	userH.RegisterRoutes(router)

	productH := handler.NewProduct(app)
	productH.RegisterRoutes(router)

	orderH := handler.NewOrder(app)
	orderH.RegisterRoutes(router)

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка при запуске HTTP сервера: %s\n", err)
	}
}
