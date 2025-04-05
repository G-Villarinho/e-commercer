package pkgs

import (
	"context"
)

type contextKey string

const (
	UserIDKey    contextKey = "user_id"
	TokenKey     contextKey = "user_token"
	UserEmailKey contextKey = "user_email"
	SessionIDKey contextKey = "session_id"
)

type RequestDataCtx interface {
	SetUserID(ctx context.Context, userID string) context.Context
	SetToken(ctx context.Context, token string) context.Context
	SetEmail(ctx context.Context, email string) context.Context
	SetSessionID(ctx context.Context, sessionID string) context.Context
	GetUserID(ctx context.Context) (string, bool)
	GetToken(ctx context.Context) (string, bool)
	GetEmail(ctx context.Context) (string, bool)
	GetSessionID(ctx context.Context) (string, bool)
}

type requestDataCtx struct {
	di           *Di
	UserIDKey    contextKey
	TokenKey     contextKey
	UserEmailKey contextKey
	SessionIDKey contextKey
}

func NewRequestInfoCtx(di *Di) (RequestDataCtx, error) {
	return &requestDataCtx{
		di:           di,
		UserIDKey:    UserIDKey,
		TokenKey:     TokenKey,
		UserEmailKey: UserEmailKey,
		SessionIDKey: SessionIDKey,
	}, nil
}

func (r *requestDataCtx) SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, r.UserIDKey, userID)
}

func (r *requestDataCtx) SetToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, r.TokenKey, token)
}

func (r *requestDataCtx) SetEmail(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, r.UserEmailKey, email)
}

func (r *requestDataCtx) SetSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, r.SessionIDKey, sessionID)
}

func (r *requestDataCtx) GetUserID(ctx context.Context) (string, bool) {
	UserID, ok := ctx.Value(r.UserIDKey).(string)
	return UserID, ok
}

func (r *requestDataCtx) GetToken(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(r.TokenKey).(string)
	return token, ok
}

func (r *requestDataCtx) GetEmail(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(r.UserEmailKey).(string)
	return email, ok
}

func (r *requestDataCtx) GetSessionID(ctx context.Context) (string, bool) {
	sessionID, ok := ctx.Value(r.SessionIDKey).(string)
	return sessionID, ok
}
