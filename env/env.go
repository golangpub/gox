package env

import (
	"github.com/gopub/log"
	"github.com/spf13/viper"
	"time"
)

func init() {
	viper.AutomaticEnv()
}

func String(key string, defaultVal string) string {
	viper.SetDefault(key, defaultVal)
	v := viper.GetString(key)
	log.Debugf("Env: %s=%s", key, v)
	return v
}

func Int(key string, defaultVal int) int {
	viper.SetDefault(key, defaultVal)
	v := viper.GetInt(key)
	log.Debugf("Env: %s=%d", key, v)
	return v
}

func Int64(key string, defaultVal int64) int64 {
	viper.SetDefault(key, defaultVal)
	v := viper.GetInt64(key)
	log.Debugf("Env: %s=%d", key, v)
	return v
}

func Int32(key string, defaultVal int32) int32 {
	viper.SetDefault(key, defaultVal)
	v := viper.GetInt32(key)
	log.Debugf("Env: %s=%d", key, v)
	return v
}

func Float64(key string, defaultVal float64) float64 {
	viper.SetDefault(key, defaultVal)
	v := viper.GetFloat64(key)
	log.Debugf("Env: %s=%f", key, v)
	return v
}

func Duration(key string, defaultVal time.Duration) time.Duration {
	viper.SetDefault(key, defaultVal)
	v := viper.GetDuration(key)
	log.Debugf("Env: %s=%v", key, v)
	return v
}

func Bool(key string, defaultVal bool) bool {
	viper.SetDefault(key, defaultVal)
	v := viper.GetBool(key)
	log.Debugf("Env: %s=%t", key, v)
	return v
}
