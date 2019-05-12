package main

import (
	handler "../internal/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"lb_account_svc/configs"
	"lb_account_svc/internal/database/psql"
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
	routerErr := router.Run(":" + conf.GetString("server.port"))
	if routerErr != nil {
		fmt.Println(routerErr)
	}
}

func initializeRoutes(config *configs.ViperConfigReader) {
	//Passing configurations
	handler.SetConfig(config)

	//Mapping endpoints
	router.GET(config.GetString("endpoints.products.my-products"), handler.ListMyProducts(db))
	router.GET(config.GetString("endpoints.products.product-transactions"), handler.ProductTransactions(db))
	router.POST(config.GetString("endpoints.beneficiary.enrollment"), handler.Enrollment(db))
	router.PATCH(config.GetString("endpoints.beneficiary.email-update"), handler.UpdateEmail(db))
	router.DELETE(config.GetString("endpoints.beneficiary.deletion"), handler.Delete(db))
}
