package log

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

// Init init config
func Init() {
	fmt.Println("init log start")
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	fmt.Println("init log end")
}
