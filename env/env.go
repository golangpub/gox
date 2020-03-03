package env

import (
	"log"
	"time"

	"github.com/gopub/gox"
	"github.com/spf13/cast"
)

type Manager interface {
	Has(key string) bool
	Get(key string) interface{}
	Set(key string, value interface{})
}

var DefaultManager Manager = NewViperManager()

func Has(key string) bool {
	return DefaultManager.Has(key)
}

func Get(key string) interface{} {
	return DefaultManager.Get(key)
}

func Set(key string, value interface{}) {
	DefaultManager.Set(key, value)
}

func String(key string, defaultValue string) string {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	return cast.ToString(DefaultManager.Get(key))
}

func MustString(key string) string {
	v := cast.ToString(DefaultManager.Get(key))
	if v == "" {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func Int(key string, defaultValue int) int {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	return cast.ToInt(DefaultManager.Get(key))
}

func MustInt(key string) int {
	v, err := cast.ToIntE(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func Int64(key string, defaultValue int64) int64 {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	return cast.ToInt64(DefaultManager.Get(key))
}

func MustInt64(key string) int64 {
	v, err := cast.ToInt64E(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func Float64(key string, defaultValue float64) float64 {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	return cast.ToFloat64(DefaultManager.Get(key))
}

func MustFloat64(key string) float64 {
	v, err := cast.ToFloat64E(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func Duration(key string, defaultValue time.Duration) time.Duration {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	return cast.ToDuration(DefaultManager.Get(key))
}

func MustDuration(key string) time.Duration {
	v, err := cast.ToDurationE(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func Bool(key string, defaultValue bool) bool {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	return cast.ToBool(DefaultManager.Get(key))
}

func MustBool(key string) bool {
	v, err := cast.ToBoolE(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func IntSlice(key string, defaultValue []int) []int {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	return cast.ToIntSlice(DefaultManager.Get(key))
}

func MustIntSlice(key string) []int {
	v, err := cast.ToIntSliceE(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	if len(v) == 0 {
		log.Panicf("%s is empty", key)
	}
	return v
}

func StringSlice(key string, defaultValue []string) []string {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	return cast.ToStringSlice(DefaultManager.Get(key))
}

func MustStringSlice(key string) []string {
	v, err := cast.ToStringSliceE(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	if len(v) == 0 {
		log.Panicf("%s is empty", key)
	}
	return v
}

func Map(key string, defaultValue gox.M) gox.M {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	return cast.ToStringMap(DefaultManager.Get(key))
}

func MustMap(key string) gox.M {
	v, err := cast.ToStringMapE(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	if len(v) == 0 {
		log.Panicf("%s is empty", key)
	}
	return v
}
