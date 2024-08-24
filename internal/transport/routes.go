package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/shiffoo/wb-nats-streaming/internal/config"
	"github.com/shiffoo/wb-nats-streaming/internal/logger"
	"github.com/shiffoo/wb-nats-streaming/internal/transport/handler"
)

func InitRouter() error {
	logg := logger.GetLogger()
	confServ := config.CONFIG.Server
	router := gin.Default()
	CreateRouter(router)

	if err := router.Run(confServ.URL); err != nil {
		logg.Panic(err)
		return err
	}
	logg.Info("Router init")
	return nil
}

func CreateRouter(r *gin.Engine) {
	r.LoadHTMLGlob("../../internal/templates/*")

	routes := r.Group("/orders")
	routes.GET("/", handler.GetAll)
	routes.POST("/", handler.AddFromPub)
	routes.GET("/:id", handler.Get)

	routes.GET("/html", handler.RenderOrdersHTML)
}
