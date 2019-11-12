package gox

import (
	"context"
)

const (
	keyDeviceID    = "device_id"
	keyRemoteAddr  = "remote"
	keyPoint       = "coordinate"
	keyAccessToken = "access_token"
	keyLoginID     = "login_id"
	keyTraceID     = "trace_id"
)

func GetLoginID(ctx context.Context) int64 {
	id, _ := ctx.Value(keyLoginID).(int64)
	return id
}

func ContextWithLoginID(ctx context.Context, id int64) context.Context {
	if id == 0 {
		return ctx
	}
	return context.WithValue(ctx, keyLoginID, id)
}

func GetAccessToken(ctx context.Context) string {
	token, _ := ctx.Value(keyAccessToken).(string)
	return token
}

func ContextWithAccessToken(ctx context.Context, token string) context.Context {
	if len(token) == 0 {
		return ctx
	}
	return context.WithValue(ctx, keyAccessToken, token)
}

func GetRemoteAddr(ctx context.Context) string {
	ip, _ := ctx.Value(keyRemoteAddr).(string)
	return ip
}

func ContextWithRemoteAddr(ctx context.Context, addr string) context.Context {
	if len(addr) == 0 {
		return ctx
	}
	return context.WithValue(ctx, keyRemoteAddr, addr)
}

func GetDeviceID(ctx context.Context) string {
	id, _ := ctx.Value(keyDeviceID).(string)
	return id
}

func ContextWithDeviceID(ctx context.Context, deviceID string) context.Context {
	if len(deviceID) == 0 {
		return ctx
	}
	return context.WithValue(ctx, keyDeviceID, deviceID)
}

func GetTraceID(ctx context.Context) string {
	id, _ := ctx.Value(keyTraceID).(string)
	return id
}

func ContextWithTraceID(ctx context.Context, traceID string) context.Context {
	if len(traceID) == 0 {
		return ctx
	}
	return context.WithValue(ctx, keyTraceID, traceID)
}

func GetPoint(ctx context.Context) *Point {
	id, _ := ctx.Value(keyPoint).(*Point)
	return id
}

func ContextWithPoint(ctx context.Context, location *Point) context.Context {
	if location == nil {
		return ctx
	}
	return context.WithValue(ctx, keyPoint, location)
}

func DetachedContext(ctx context.Context) context.Context {
	newCtx := context.Background()
	if token := GetAccessToken(ctx); len(token) > 0 {
		newCtx = ContextWithAccessToken(newCtx, token)
	}
	if deviceID := GetDeviceID(ctx); len(deviceID) > 0 {
		newCtx = ContextWithDeviceID(newCtx, deviceID)
	}
	if c := GetPoint(ctx); c != nil {
		newCtx = ContextWithPoint(newCtx, c)
	}
	if addr := GetRemoteAddr(ctx); len(addr) > 0 {
		newCtx = ContextWithRemoteAddr(newCtx, addr)
	}
	if traceID := GetTraceID(ctx); len(traceID) > 0 {
		newCtx = ContextWithTraceID(newCtx, traceID)
	}
	if loginID := GetLoginID(ctx); loginID > 0 {
		newCtx = ContextWithLoginID(newCtx, loginID)
	}
	return newCtx
}
