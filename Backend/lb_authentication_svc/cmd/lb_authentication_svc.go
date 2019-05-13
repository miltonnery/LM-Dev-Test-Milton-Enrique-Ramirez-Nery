package main

import (
	handler "../internal/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"lb_authentication_svc/configs"
	"lb_authentication_svc/internal/database/psql"
)

var router *gin.Engine
var db *gorm.DB
var err error
var Profile *string

func main() {

	//Getting YAML configurations
	conf := configs.ConfReader

	db, err = psql.GetPostgreSQLDB(conf)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

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
	router.POST(config.GetString("endpoints.login"), handler.Login(db))
}
