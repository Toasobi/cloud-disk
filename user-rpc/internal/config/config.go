package config

import (
	"os"

	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("user")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/etc")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
