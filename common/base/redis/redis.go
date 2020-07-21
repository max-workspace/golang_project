package redis

import (
	"time"
)

// Instance redis interface
type Instance interface {
	Close() error
	Exists(keys ...string) error
	Expire(key string, expiration time.Duration) error
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Del(keys ...string) error
}
