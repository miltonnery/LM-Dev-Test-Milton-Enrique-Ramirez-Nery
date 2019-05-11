package configs

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"lb_authentication_svc/internal/model/config"
	"log"
	"time"
)

type Reader interface {
	GetAllKeys() []string
	Get(key string) interface{}
	GetBool(key string) bool
	GetString(key string) string
}

type ViperConfigReader struct {
	viper *viper.Viper
}

var ConfReader *ViperConfigReader

func (v ViperConfigReader) GetAllKeys() []string{
	return v.viper.AllKeys()
}

func (v ViperConfigReader) Get(key string) interface{} {
	return v.viper.Get(key)
}

func (v ViperConfigReader) GetBool(key string) bool {
	return v.viper.GetBool(key)
}

func (v ViperConfigReader) GetString(key string) string {
	return v.viper.GetString(key)
}


func init() {

	profile := flag.String("profile", "application-local", "The path to application configurations")
	flag.Parse()
	v:= viper.New()
	v.SetConfigName(*profile)
	v.SetConfigType("yml")
	v.AddConfigPath("./configs/environment/")
	v.SetDefault("server.port", "9090")
	//Read it

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Not able to read configuration, %s", err)
	}

	ConfReader = &ViperConfigReader{
		viper: v,
	}
	var configs model.Configurations


	err := v.Unmarshal(&configs)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}


	go func() {
		for {
			time.Sleep(time.Second * 5)
			v.WatchConfig()
			v.OnConfigChange(func(e fsnotify.Event) {
				log.Println("config file changed", e.Name)
			})
		}
	}()
}