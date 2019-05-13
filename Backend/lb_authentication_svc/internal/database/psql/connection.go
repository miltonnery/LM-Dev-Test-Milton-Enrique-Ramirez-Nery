package psql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"lb_authentication_svc/configs"
	"strconv"
	"time"
)

func GetPostgreSQLDB(cfg *configs.ViperConfigReader) (db *gorm.DB, err error) {
	connectionString :=
			"host="      + cfg.Get("datasource.postgres.host").(string) +
			" port="     + strconv.Itoa(cfg.Get("datasource.postgres.port").(int)) +
			" user="     + cfg.Get("datasource.postgres.user").(string) +
			" dbname="   + cfg.Get("datasource.postgres.dbname").(string) +
			" password=" + cfg.Get("datasource.postgres.pass").(string) +
			" sslmode="  + cfg.Get("datasource.postgres.sslmode").(string)

	db, err = gorm.Open(cfg.Get("datasource.postgres.driver").(string), connectionString)

	if err != nil {
		logrus.Error("CONNECTION ERROR! -> %v", err)
	}

	//Configuring connection pool
	maxOpenConns := cfg.Get("connection.maxOpenConns").(int)
	maxIdleConns := cfg.Get("connection.maxIdleCons").(int)
	maxConnLifetime := cfg.Get("connection.maxConnLifetime").(int)

	db.DB().SetMaxOpenConns(maxOpenConns)
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetConnMaxLifetime(time.Duration(maxConnLifetime) * time.Minute)

	return
}
