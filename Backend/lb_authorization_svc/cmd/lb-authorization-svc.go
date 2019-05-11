package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"lb_authorization_svc/configs"
	"lb_authorization_svc/internal/middleware/gateway"
)

var db *gorm.DB
var err error
var Profile *string

func main() {
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function
	conf := configs.ConfReader

	//Activating the Catch-All Routing
	engine := gin.Default()
	engine.NoRoute(gateway.Redirector(conf))
	engine.Run(":" + conf.GetString("server.port"))
}
