package gox

import "time"

type ConfigReader interface {
	GetString(key string) string
	GetInt64(key string) string
	GetBool(key string) bool
	GetDuration(key string) time.Duration
}
