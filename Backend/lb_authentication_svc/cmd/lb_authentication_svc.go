package main

import (
	handler "../internal/middleware"
	"github.com/gin-gonic/gin"
	"lb_authentication_svc/configs"
)

var router *gin.Engine

func main() {

	//Getting YAML configurations
	conf := configs.ConfReader

	//Initializing Gin Gonic
	router = gin.Default()
	initializeRoutes(conf)

	//Starting server
	router.Run(":" + conf.GetString("server.port"))
}

func initializeRoutes(config *configs.ViperConfigReader) {
	//Passing configurations
	handler.SetConfig(config)

	//Mapping endpoints
	router.POST(config.GetString("endpoints.login"), handler.Login)
}
