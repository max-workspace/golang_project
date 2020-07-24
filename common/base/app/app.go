package app

import (
	"project/common/base/config"
	"project/common/base/log"
	"project/common/base/log/zapadapter"
	"project/common/base/redis"
	"project/common/base/redis/goredisadapter"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type application struct {
	config config.Instance
	log    log.Instance
	redis  redis.Instance
}

type logEntry struct {
	log   log.Instance
	ready chan struct{} // closed when log is ready
}

var (
	app            *application
	loadAppOnce    sync.Once
	loadConfigOnce sync.Once
	loadLogOnce    sync.Once
	loadRedisOnce  sync.Once

	logMap         map[string]*logEntry
	logMapLock     sync.Mutex
	loadLogMapOnce sync.Once
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
		logInstance := zapadapter.New(app.GetConfig().GetString("app.log.path.debug"), zapcore.DebugLevel, 128, 30, 7, true, app.GetConfig().GetString("app.name"))
		app.log = logInstance
	})
	return app.log
}

// GetRedis get redis singleton
func (app *application) GetRedis() redis.Instance {
	loadRedisOnce.Do(func() {
		redisInstance := goredisadapter.New(app.GetConfig().GetString("app.redis.localDemo.addr"), app.GetConfig().GetString("app.redis.localDemo.password"), app.GetConfig().GetInt("app.redis.localDemo.DB"))
		app.redis = redisInstance
	})
	return app.redis
}

// GetCustomLog get custom log
func (app *application) GetCustomLog(key, path, project string) log.Instance {
	loadLogMapOnce.Do(func() {
		logMap = make(map[string]*logEntry)
	})
	logMapLock.Lock()
	item := logMap[key]
	if item == nil {
		item = &logEntry{ready: make(chan struct{})}
		logMap[key] = item
		logMapLock.Unlock()
		item.log = zapadapter.NewCustom(key, path, zapcore.DebugLevel, 128, 30, 7, true, project)
		close(item.ready)
	} else {
		logMapLock.Unlock()
		<-item.ready
	}
	return item.log
}
