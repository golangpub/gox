package gox

import (
	"context"
)

const (
	keyDeviceID    = "_device_id"
	keyRemoteAddr  = "_remote"
	keyLocation    = "_location"
	keyAccessToken = "_access_token"
	keyUserID      = "_user_id"
	keyTraceID     = "_trace_id"
	keyUser        = "_user"
)

func GetUserID(ctx context.Context) int64 {
	id, _ := ctx.Value(keyUserID).(int64)
	return id
}

func WithUserID(ctx context.Context, id int64) context.Context {
	if id == 0 {
		return ctx
	}
	return context.WithValue(ctx, keyUserID, id)
}

func GetUser(ctx context.Context) interface{} {
	return ctx.Value(keyUser)
}

func WithUser(ctx context.Context, u interface{}) context.Context {
	return context.WithValue(ctx, keyUser, u)
}

func GetAccessToken(ctx context.Context) string {
	token, _ := ctx.Value(keyAccessToken).(string)
	return token
}

func WithAccessToken(ctx context.Context, token string) context.Context {
	if len(token) == 0 {
		return ctx
	}
	return context.WithValue(ctx, keyAccessToken, token)
}

func GetRemoteAddr(ctx context.Context) string {
	ip, _ := ctx.Value(keyRemoteAddr).(string)
	return ip
}

func WithRemoteAddr(ctx context.Context, addr string) context.Context {
	if len(addr) == 0 {
		return ctx
	}
	return context.WithValue(ctx, keyRemoteAddr, addr)
}

func GetDeviceID(ctx context.Context) string {
	id, _ := ctx.Value(keyDeviceID).(string)
	return id
}

func WithDeviceID(ctx context.Context, deviceID string) context.Context {
	if len(deviceID) == 0 {
		return ctx
	}
	return context.WithValue(ctx, keyDeviceID, deviceID)
}

func GetTraceID(ctx context.Context) string {
	id, _ := ctx.Value(keyTraceID).(string)
	return id
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	if len(traceID) == 0 {
		return ctx
	}
	return context.WithValue(ctx, keyTraceID, traceID)
}

func GetLocation(ctx context.Context) *Point {
	id, _ := ctx.Value(keyLocation).(*Point)
	return id
}

func WithLocation(ctx context.Context, location *Point) context.Context {
	if location == nil {
		return ctx
	}
	return context.WithValue(ctx, keyLocation, location)
}

func DetachedContext(ctx context.Context) context.Context {
	newCtx := context.Background()
	if token := GetAccessToken(ctx); len(token) > 0 {
		newCtx = WithAccessToken(newCtx, token)
	}
	if deviceID := GetDeviceID(ctx); len(deviceID) > 0 {
		newCtx = WithDeviceID(newCtx, deviceID)
	}
	if c := GetLocation(ctx); c != nil {
		newCtx = WithLocation(newCtx, c)
	}
	if addr := GetRemoteAddr(ctx); len(addr) > 0 {
		newCtx = WithRemoteAddr(newCtx, addr)
	}
	if traceID := GetTraceID(ctx); len(traceID) > 0 {
		newCtx = WithTraceID(newCtx, traceID)
	}
	if loginID := GetUserID(ctx); loginID > 0 {
		newCtx = WithUserID(newCtx, loginID)
	}
	return newCtx
}
