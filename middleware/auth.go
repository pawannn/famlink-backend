package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pawannn/famlink/api"
	port "github.com/pawannn/famlink/port/token"
)

func Auth(tokenService *port.TokenPort) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			api.SendResponse(c, http.StatusBadRequest, "missing `Authorization` token", nil)
			c.Abort()
			return
		}
		userID, err := tokenService.ParseUserToken(authToken)
		if err != nil {
			api.SendResponse(c, http.StatusBadGateway, "Invalid or expired token", nil)
			c.Abort()
			return
		}
		userCtx := UserContext{User_id: userID}
		ctx := AttachContext(c.Request.Context(), userCtx)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
