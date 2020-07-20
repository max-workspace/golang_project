package main

import (
	"fmt"
	"project/common/config"
	"project/common/log"

	"github.com/spf13/viper"
)

func main() {
	config.Init()
	log.Init()

	msg := fmt.Sprintf("project init finish! project name: %s", viper.GetString("app.name"))
	log.DebugLogger.Debug(msg)
}
