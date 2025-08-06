package api

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Code          int    `json:"code"`
	ClientMessage string `json:"clientMessage"`
	Data          any    `json:"data"`
}

func SendResponse(c *gin.Context, code int, clientMessage string, data any) {
	errResponse := ApiResponse{
		Code:          code,
		ClientMessage: clientMessage,
		Data:          data,
	}
	c.JSON(code, errResponse)
}
