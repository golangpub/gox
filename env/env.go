package env

import (
	"github.com/spf13/viper"
	"time"
)

func init() {
	viper.AutomaticEnv()
}

func String(key string, defaultVal string) string {
	viper.SetDefault(key, defaultVal)
	return viper.GetString(key)
}

func Int(key string, defaultVal int) int {
	viper.SetDefault(key, defaultVal)
	return viper.GetInt(key)
}

func Int64(key string, defaultVal int64) int64 {
	viper.SetDefault(key, defaultVal)
	return viper.GetInt64(key)
}

func Int32(key string, defaultVal int32) int32 {
	viper.SetDefault(key, defaultVal)
	return viper.GetInt32(key)
}

func Float64(key string, defaultVal float64) float64 {
	viper.SetDefault(key, defaultVal)
	return viper.GetFloat64(key)
}

func Duration(key string, defaultVal time.Duration) time.Duration {
	viper.SetDefault(key, defaultVal)
	return viper.GetDuration(key)
}

func Bool(key string, defaultVal bool) bool {
	viper.SetDefault(key, defaultVal)
	return viper.GetBool(key)
}
