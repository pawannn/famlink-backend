package middleware

import (
	"context"

	"github.com/pawannn/famly/pkg/constants"
)

type UserContext struct {
	User_id string `json:"user_id"`
}

func AttachContext(ctx context.Context, userContext UserContext) context.Context {
	return context.WithValue(ctx, constants.USER_CONTEXT_KEY, userContext)
}

func GetUserContext(ctx context.Context) (UserContext, bool) {
	val := ctx.Value(constants.USER_CONTEXT_KEY)
	userCtx, ok := val.(UserContext)
	return userCtx, ok
}
