package app

import (
	"project/common/base/config"
	"project/common/base/log"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type application struct {
	config config.Instance
	log    log.Instance
}

var (
	app            *application
	loadAppOnce    sync.Once
	loadConfigOnce sync.Once
	loadLogOnce    sync.Once
)

// Instance get application singleton
func Instance() *application {
	loadAppOnce.Do(func() {
		app = &application{}
	})
	return app
}

// GetConfig get config singleton
func (app *application) GetConfig() config.Instance {
	loadConfigOnce.Do(func() {
		configInstance := viper.New()
		configInstance.SetConfigName("service")
		configInstance.AddConfigPath("./config")
		err := configInstance.ReadInConfig()
		if err != nil {
			panic(err)
		}
		app.config = configInstance
	})
	return app.config
}

// GetLog get log singleton
func (app *application) GetLog() log.Instance {
	loadLogOnce.Do(func() {
		logInstance := log.NewLogger(app.GetConfig().GetString("app.log.path.debug"), zapcore.DebugLevel, 128, 30, 7, true, app.GetConfig().GetString("app.name"))
		app.log = logInstance
	})
	return app.log
}
