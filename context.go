package gox

import (
	"context"
)

const (
	keyDeviceID    = "device_id"
	keyRemoteAddr  = "remote"
	keyCoordinate  = "coordinate"
	keyAccessToken = "access_token"
	keyLoginID     = "login_id"
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

func ContextWithDeviceID(ctx context.Context, deviceId string) context.Context {
	if len(deviceId) == 0 {
		return ctx
	}
	return context.WithValue(ctx, keyDeviceID, deviceId)
}

func GetCoordinate(ctx context.Context) *Coordinate {
	id, _ := ctx.Value(keyCoordinate).(*Coordinate)
	return id
}

func ContextWithCoordinate(ctx context.Context, location *Coordinate) context.Context {
	if location == nil {
		return ctx
	}
	return context.WithValue(ctx, keyCoordinate, location)
}

func DetachedContext(ctx context.Context) context.Context {
	newCtx := context.Background()
	if token := GetAccessToken(ctx); len(token) > 0 {
		newCtx = ContextWithAccessToken(newCtx, token)
	}
	if deviceID := GetDeviceID(ctx); len(deviceID) > 0 {
		newCtx = ContextWithDeviceID(newCtx, deviceID)
	}
	if c := GetCoordinate(ctx); c != nil {
		newCtx = ContextWithCoordinate(newCtx, c)
	}
	if addr := GetRemoteAddr(ctx); len(addr) > 0 {
		newCtx = ContextWithRemoteAddr(newCtx, addr)
	}
	return newCtx
}
