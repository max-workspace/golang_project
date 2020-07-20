package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Init init config
func Init() {
	fmt.Println("init config begin...")
	viper.SetConfigName("service")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("init config done")
}
