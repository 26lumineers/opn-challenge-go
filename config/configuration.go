package config

import (
	"fmt"
	"opn/model"
	"strings"

	"github.com/spf13/viper"
)

func SetUpConfiguration(config *model.Configuration) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("fatal error config file:%s\n", err))
	}

	err = viper.Unmarshal(config)
	if err!=nil{
		panic(fmt.Sprintf("unmarshal config failed :: %s",err))
	}

}
