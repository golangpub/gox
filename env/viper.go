package env

import (
	"time"

	"github.com/gopub/log"
	"github.com/spf13/viper"
)

type ViperManager struct {
}

func NewViperManager() *ViperManager {
	viper.AutomaticEnv()
	return &ViperManager{}
}

func (m *ViperManager) Has(key string) bool {
	return viper.IsSet(key)
}

func (m *ViperManager) SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

func (m *ViperManager) String(key string, defaultVal string) string {
	v := defaultVal
	if viper.IsSet(key) {
		log.Debugf("Env: missing %s", key)
		v = viper.GetString(key)
	}
	log.Debugf("Env: %s=%s", key, v)
	return v
}

func (m *ViperManager) Int(key string, defaultVal int) int {
	v := defaultVal
	if viper.IsSet(key) {
		log.Debugf("Env: missing %s", key)
		v = viper.GetInt(key)
	}
	log.Debugf("Env: %s=%d", key, v)
	return v
}

func (m *ViperManager) Int64(key string, defaultVal int64) int64 {
	v := defaultVal
	if viper.IsSet(key) {
		log.Debugf("Env: missing %s", key)
		v = viper.GetInt64(key)
	}
	log.Debugf("Env: %s=%d", key, v)
	return v
}

func (m *ViperManager) Float64(key string, defaultVal float64) float64 {
	v := defaultVal
	if viper.IsSet(key) {
		log.Debugf("Env: missing %s", key)
		v = viper.GetFloat64(key)
	}
	log.Debugf("Env: %s=%f", key, v)
	return v
}

func (m *ViperManager) Duration(key string, defaultVal time.Duration) time.Duration {
	v := defaultVal
	if viper.IsSet(key) {
		log.Debugf("Env: missing %s", key)
		v = viper.GetDuration(key)
	}
	log.Debugf("Env: %s=%v", key, v)
	return v
}

func (m *ViperManager) Bool(key string, defaultVal bool) bool {
	v := defaultVal
	if viper.IsSet(key) {
		log.Debugf("Env: missing %s", key)
		v = viper.GetBool(key)
	}
	log.Debugf("Env: %s=%t", key, v)
	return v
}

func (m *ViperManager) IntSlice(key string, defaultVal []int) []int {
	v := defaultVal
	if viper.IsSet(key) {
		log.Debugf("Env: missing %s", key)
		v = viper.GetIntSlice(key)
	}
	log.Debugf("Env: %s=%v", key, v)
	return v
}

func (m *ViperManager) StringSlice(key string, defaultVal []string) []string {
	v := defaultVal
	if viper.IsSet(key) {
		log.Debugf("Env: missing %s", key)
		v = viper.GetStringSlice(key)
	}
	log.Debugf("Env: %s=%v", key, v)
	return v
}
