package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pawannn/famlink/api"
	domain "github.com/pawannn/famlink/domain/users"
)

func (u *User) VerifyPhone(c *gin.Context) {
}

func (u *User) RegisterUser(c *gin.Context) {
	var userDetails domain.UserSchema
	if err := c.BindJSON(&userDetails); err != nil {
		api.SendResponse(c, http.StatusInternalServerError, "Unable to read body", nil)
		return
	}
	user, err := u.UserService.Register(userDetails.Name, userDetails.Phone, userDetails.Country)
	if err != nil {
		api.SendResponse(c, http.StatusInternalServerError, "Unable to register user", err.Error())
		return
	}
	api.SendResponse(c, http.StatusOK, "User registered successfully", user)
}
