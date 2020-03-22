package gox

import (
	"context"
	"github.com/gopub/types"
)

type contextKey int

const (
	ckDeviceID contextKey = iota + 1
	ckRemoteAddr
	ckLocation
	ckAccessToken
	ckUserID
	ckTraceID
	ckUser
)

func GetUserID(ctx context.Context) int64 {
	id, _ := ctx.Value(ckUserID).(int64)
	return id
}

func WithUserID(ctx context.Context, id int64) context.Context {
	if id == 0 {
		return ctx
	}
	return context.WithValue(ctx, ckUserID, id)
}

func GetUser(ctx context.Context) interface{} {
	return ctx.Value(ckUser)
}

func WithUser(ctx context.Context, u interface{}) context.Context {
	return context.WithValue(ctx, ckUser, u)
}

func GetAccessToken(ctx context.Context) string {
	token, _ := ctx.Value(ckAccessToken).(string)
	return token
}

func WithAccessToken(ctx context.Context, token string) context.Context {
	if len(token) == 0 {
		return ctx
	}
	return context.WithValue(ctx, ckAccessToken, token)
}

func GetRemoteAddr(ctx context.Context) string {
	ip, _ := ctx.Value(ckRemoteAddr).(string)
	return ip
}

func WithRemoteAddr(ctx context.Context, addr string) context.Context {
	if len(addr) == 0 {
		return ctx
	}
	return context.WithValue(ctx, ckRemoteAddr, addr)
}

func GetDeviceID(ctx context.Context) string {
	id, _ := ctx.Value(ckDeviceID).(string)
	return id
}

func WithDeviceID(ctx context.Context, deviceID string) context.Context {
	if len(deviceID) == 0 {
		return ctx
	}
	return context.WithValue(ctx, ckDeviceID, deviceID)
}

func GetTraceID(ctx context.Context) string {
	id, _ := ctx.Value(ckTraceID).(string)
	return id
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	if len(traceID) == 0 {
		return ctx
	}
	return context.WithValue(ctx, ckTraceID, traceID)
}

func GetLocation(ctx context.Context) *types.Point {
	id, _ := ctx.Value(ckLocation).(*types.Point)
	return id
}

func WithLocation(ctx context.Context, location *types.Point) context.Context {
	if location == nil {
		return ctx
	}
	return context.WithValue(ctx, ckLocation, location)
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
