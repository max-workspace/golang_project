package main

import (
	"fmt"
	"project/common/config"
	"project/common/log"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("start project")
	config.Init()
	log.Init()
	fmt.Printf("project name: %s\n", viper.GetString("app.name"))
}
