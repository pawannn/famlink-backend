package user

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nyaruka/phonenumbers"
	"github.com/pawannn/famlink/api"
	domain "github.com/pawannn/famlink/domain/users"
	"github.com/pawannn/famlink/pkg"
	"github.com/pawannn/famlink/pkg/constants"
)

func (u *User) RegisterUser(c *gin.Context) {
	var userDetails domain.UserSchema
	if err := c.BindJSON(&userDetails); err != nil {
		api.SendResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	name := strings.TrimSpace(userDetails.Name)
	if name == "" {
		api.SendResponse(c, http.StatusBadRequest, "Name cannot be empty", nil)
		return
	}
	countryCode := strings.ToUpper(strings.TrimSpace(userDetails.Country))
	if !pkg.ValidateCountry(countryCode) {
		api.SendResponse(c, http.StatusBadRequest, "Invalid country code", nil)
		return
	}
	parsedPhone, err := phonenumbers.Parse(userDetails.Phone, countryCode)
	if err != nil || !phonenumbers.IsValidNumber(parsedPhone) {
		api.SendResponse(c, http.StatusBadRequest, "Invalid phone number for the specified country", nil)
		return
	}
	formattedPhone := phonenumbers.Format(parsedPhone, phonenumbers.E164)
	user, err := u.UserService.Register(name, formattedPhone, countryCode)
	if err != nil {
		if strings.Contains(err.Error(), constants.ERR_USER_EXIST) {
			api.SendResponse(c, http.StatusConflict, "User already exists", nil)
			return
		}
		api.SendResponse(c, http.StatusInternalServerError, "Unable to register user", err)
		return
	}

	api.SendResponse(c, http.StatusOK, "User registered successfully", user)
}
