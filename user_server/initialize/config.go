package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go-servers/user_server/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	//debug := GetEnvInfo("mistyu")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user_server/%s.yaml", configFilePrefix)
	//if debug {
	//	configFileName = fmt.Sprintf("user_server/%s-debug.yaml", configFilePrefix)
	//}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
}
