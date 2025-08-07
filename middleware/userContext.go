package middleware

import "context"

type ContextKey string

var UserContextKey ContextKey = "UserContext"

type UserContext struct {
	User_id string `json:"user_id"`
}

func AttachContext(ctx context.Context, userContext UserContext) context.Context {
	return context.WithValue(ctx, UserContextKey, userContext)
}

func GetUserContext(ctx context.Context) (UserContext, bool) {
	val := ctx.Value(UserContextKey)
	userCtx, ok := val.(UserContext)
	return userCtx, ok
}
