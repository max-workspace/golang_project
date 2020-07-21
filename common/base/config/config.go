package config

import (
	"time"
)

// Instance config interface
type Instance interface {
	IsSet(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetIntSlice(key string) []int
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetUint(key string) uint
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	GetFloat64(key string) float64
	GetBool(key string) bool
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
}
