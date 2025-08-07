package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	token "github.com/pawannn/famlink/adapter/token/jwt"
	"github.com/pawannn/famlink/api"
)

func Auth(tokenService token.TokenRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			api.SendResponse(c, http.StatusBadRequest, "missing `Authorization` token", nil)
			c.Abort()
			return
		}
		claims, err := tokenService.ValidateJWT(authToken)
		if err != nil {
			api.SendResponse(c, http.StatusBadGateway, "Invalid or expired token", nil)
			c.Abort()
			return
		}
		userCtx := UserContext{User_id: claims.UserID}
		ctx := AttachContext(c.Request.Context(), userCtx)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
