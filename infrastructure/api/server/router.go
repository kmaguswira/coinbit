package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/coinbit/application/config"
	"github.com/kmaguswira/coinbit/infrastructure/api/controllers"
	"github.com/kmaguswira/coinbit/infrastructure/api/middlewares"
)

func SetupRouter() *gin.Engine {
	if config.GetConfig().Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	loggerMiddleware := middlewares.NewLoggerMiddleware()
	corsMiddleware := middlewares.NewCorsMiddleware()
	idempotentMiddleware := middlewares.NewIdempotentMiddleware()
	errorMiddleware := middlewares.NewErrorMiddleware()

	router.Use(gin.Recovery())
	router.Use(loggerMiddleware.GetHandler())
	router.Use(corsMiddleware.GetHandler())
	router.Use(idempotentMiddleware.GetHandler())
	router.Use(errorMiddleware.GetHandler())

	RouterV1(router)

	return router
}

func RouterV1(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		healthGroup := v1.Group("health")
		{
			health := controllers.NewHealthController()
			healthGroup.GET("/check", health.Check)
		}

		coinbitGroup := v1.Group("coinbit")
		{
			coinbit := controllers.NewCoinbitController()
			coinbitGroup.POST("/deposit", coinbit.Deposit)
			coinbitGroup.GET("/balance/:id", coinbit.GetBalance)
		}
	}
}
