package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-house/api"
	"github.com/go-house/library"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	initConf() // 读取配置
	api.Router(&r.RouterGroup)
	r.Run()
}

func initConf() *viper.Viper {
	viper.SetConfigName("dev")
	viper.AddConfigPath("conf/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	library.InitMysql(viper.GetStringMapString("db_house"))

	library.InitWxSDK(viper.GetStringMapString("wxsdk"))
	return viper.GetViper()
}
