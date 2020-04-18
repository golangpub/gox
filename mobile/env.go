package mobile

import "github.com/golangpub/gox/env"

func EnvString(key string) string {
	return env.String(key, "")
}

func EnvInt(key string) int {
	return env.Int(key, 0)
}

func EnvFloat64(key string) float64 {
	return env.Float64(key, 0)
}
