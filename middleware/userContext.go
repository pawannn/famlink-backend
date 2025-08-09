package middleware

import (
	"context"

	"github.com/pawannn/famlink/pkg/constants"
)

type UserContext struct {
	User_id string `json:"user_id"`
}

func AttachContext(ctx context.Context, userContext UserContext) context.Context {
	return context.WithValue(ctx, constants.UserContextKey, userContext)
}

func GetUserContext(ctx context.Context) (UserContext, bool) {
	val := ctx.Value(constants.UserContextKey)
	userCtx, ok := val.(UserContext)
	return userCtx, ok
}
