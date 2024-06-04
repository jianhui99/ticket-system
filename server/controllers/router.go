package controller

import (
	"ticket-system/controllers/home"
	"ticket-system/controllers/queue"
	"ticket-system/middlewares"

	"github.com/gin-gonic/gin"
)

func route(engine *gin.Engine) {
	engine.Use(middlewares.CorsMiddleware)
	rg := engine.Group("/api/v1")
	{
		rg.GET("/home", home.HandleGetHome)

		rg.POST("/queues", queue.HandleInitQueue)
	}
}
