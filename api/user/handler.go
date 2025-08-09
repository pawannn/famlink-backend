package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nyaruka/phonenumbers"
	"github.com/pawannn/famlink/api"
	domain "github.com/pawannn/famlink/core/domain/users"
	"github.com/pawannn/famlink/middleware"
	"github.com/pawannn/famlink/pkg/constants"
	"github.com/pawannn/famlink/pkg/helpers"
)

func (u *User) ValidatePhone(c *gin.Context) {
	var userDetails domain.UserSchema
	if err := c.BindJSON(&userDetails); err != nil {
		api.SendResponse(c, http.StatusBadRequest, constants.ERR_INVALID_BODY, err.Error())
		return
	}
	countryCode := strings.ToUpper(strings.TrimSpace(userDetails.Country))
	if !helpers.ValidateCountry(countryCode) {
		api.SendResponse(c, http.StatusBadRequest, constants.ERR_INVALID_COUNTRY_CODE, nil)
		return
	}
	parsedPhone, err := phonenumbers.Parse(userDetails.Phone, countryCode)
	if err != nil || !phonenumbers.IsValidNumber(parsedPhone) {
		api.SendResponse(c, http.StatusBadRequest, constants.ERR_INVALID_PHONE, err.Error())
		return
	}
	formattedPhone := phonenumbers.Format(parsedPhone, phonenumbers.E164)
	fmt.Println(formattedPhone)
	err = u.UserSmsRepo.SendUserOTP(formattedPhone)
	if err != nil {
		api.SendResponse(c, http.StatusInternalServerError, "unable to send OTP", err.Error())
		return
	}
	api.SendResponse(c, http.StatusOK, "OTP sent to phone number", nil)
}

func (u *User) VerifyPhone(c *gin.Context) {
	var payload domain.VerifyPhonePayload
	if err := c.BindJSON(&payload); err != nil {
		api.SendResponse(c, http.StatusBadRequest, constants.ERR_INVALID_BODY, err.Error())
		return
	}

	countryCode := strings.ToUpper(strings.TrimSpace(payload.Country))
	if !helpers.ValidateCountry(countryCode) {
		api.SendResponse(c, http.StatusBadRequest, constants.ERR_INVALID_COUNTRY_CODE, nil)
		return
	}

	parsedPhone, err := phonenumbers.Parse(payload.Phone, countryCode)
	if err != nil || !phonenumbers.IsValidNumber(parsedPhone) {
		api.SendResponse(c, http.StatusBadRequest, constants.ERR_INVALID_PHONE, err.Error())
		return
	}
	formattedPhone := phonenumbers.Format(parsedPhone, phonenumbers.E164)
	user, err := u.UserRepo.GetUserByPhone(formattedPhone)
	if err != nil {
		api.SendResponse(c, http.StatusInternalServerError, constants.ERR_FETCH_USER, err.Error())
		return
	}
	ok, err := u.UserSmsRepo.VerifyUserOTP(formattedPhone, payload.OTP)
	if err != nil {
		api.SendResponse(c, http.StatusInternalServerError, "unable to verify OTP", err.Error())
		return
	}
	if !ok {
		api.SendResponse(c, http.StatusUnauthorized, "Incorrect OTP", nil)
		return
	}

	if user == nil {
		random := helpers.GenerateRandom()
		username := fmt.Sprintf("user-%d", random)
		userDetails := domain.UserSchema{
			ID:      helpers.GenerateUUID(),
			Name:    username,
			Phone:   formattedPhone,
			Avatar:  "",
			Country: countryCode,
		}
		user, err = u.UserRepo.Register(userDetails)
		if err != nil {
			if strings.Contains(err.Error(), constants.ERR_USER_EXIST) {
				api.SendResponse(c, http.StatusConflict, constants.ERR_USER_EXIST, nil)
				return
			}
			api.SendResponse(c, http.StatusInternalServerError, "Unable to register user", err.Error())
			return
		}
	}
	token, err := u.FE.Token.GenerateUserToken(user.ID)
	if err != nil {
		api.SendResponse(c, http.StatusInternalServerError, "Failed to generate token", err.Error())
		return
	}
	response := domain.VerifyPhoneResponse{
		User:  *user,
		Token: token,
	}
	api.SendResponse(c, http.StatusOK, "User phone verified successfully", response)
}

func (u *User) UpdateUser(c *gin.Context) {
	userCtx, ok := middleware.GetUserContext(c.Request.Context())
	if !ok {
		api.SendResponse(c, http.StatusUnauthorized, "unable to get userID", nil)
		return
	}
	userID := userCtx.User_id
	username := c.PostForm("username")
	_, fileErr := c.FormFile("avatar")
	if strings.TrimSpace(username) == "" && fileErr != nil {
		api.SendResponse(c, http.StatusBadRequest, "No field to update", fileErr.Error())
		return
	}
	var name *string
	var avatarURL *string
	if fileErr == nil {
		// TODO: Upload to S3
		s3Url := "https://somelink.com"
		avatarURL = &s3Url
	}
	if username != "" {
		name = &username
	}
	user, err := u.UserRepo.UpdateUser(userID, name, avatarURL)
	if err != nil {
		api.SendResponse(c, http.StatusInternalServerError, "Unable to updated user details", err.Error())
		return
	}
	err = u.UserCacheRepo.SaveUser(userID, *user)
	if err != nil {
		fmt.Printf("unable to set data to cache: %s\n", err.Error())
	}
	api.SendResponse(c, http.StatusOK, "User details updated successfully", user)
}

func (u *User) GetUser(c *gin.Context) {
	userCtx, ok := middleware.GetUserContext(c.Request.Context())
	if !ok {
		api.SendResponse(c, http.StatusUnauthorized, "unable to get userID", nil)
		return
	}
	userID := userCtx.User_id
	userCache, err := u.UserCacheRepo.GetUser(userID)
	if err != nil {
		fmt.Printf("unable to get data from cache: %s\n", err.Error())
	}
	if userCache != nil {
		api.SendResponse(c, http.StatusOK, "Successfully fetched user details", userCache)
	}

	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		api.SendResponse(c, http.StatusInternalServerError, constants.ERR_FETCH_USER, err.Error())
		return
	}
	if user == nil {
		api.SendResponse(c, http.StatusNotFound, "User is not registered", nil)
		return
	}
	err = u.UserCacheRepo.SaveUser(userID, *user)
	if err != nil {
		fmt.Printf("unable to set data to cache: %s\n", err.Error())
	}
	api.SendResponse(c, http.StatusOK, "Successfully fetched user details", user)
}
