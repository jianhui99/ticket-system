package controller

import (
	"strconv"
	"ticket-system/config"

	"github.com/gin-gonic/gin"
)

func Init() {
	engine := gin.New()
	isProd := config.Conf.App.Env == "production"
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		engine.Use(gin.Logger())
	}
	engine.Use(gin.Recovery())
	route(engine)

	port := strconv.Itoa(config.Conf.Port)
	err := engine.Run(":" + port)

	if err != nil {
		panic(err)
	}
}
