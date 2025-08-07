package user

import (
	"github.com/gin-gonic/gin"
	domain "github.com/pawannn/famlink/domain/users"
	"github.com/pawannn/famlink/middleware"
	"github.com/pawannn/famlink/pkg"
	port "github.com/pawannn/famlink/port/repository"
)

type User struct {
	FE       pkg.FamLinkEngine
	UserRepo *port.UserRepository
}

func InitUserService(fE pkg.FamLinkEngine, userService domain.UserService) *User {
	userRepo := port.InitUserService(userService)
	return &User{
		FE:       fE,
		UserRepo: userRepo,
	}
}

func (u *User) InitUserRoutes() {
	u.FE.AddRoute([]pkg.FamLinkRoute{
		{
			Route:       "/user/phone/validate",
			Method:      "POST",
			Controller:  u.ValidatePhone,
			Description: "This endpoint validates the phone and country with OTP",
			Middleware:  nil,
		},
		{
			Route:       "/user/phone/verify",
			Method:      "PUT",
			Controller:  u.VerifyPhone,
			Description: "This endpoint verify the phone with the given OTP",
			Middleware:  nil,
		},
		{
			Route:       "/user",
			Method:      "PATCH",
			Controller:  u.UpdateUser,
			Description: "This endpoint updates the username of the user",
			Middleware: []gin.HandlerFunc{
				middleware.Auth(u.FE.TokenService),
			},
		},
		{
			Route:       "/user",
			Method:      "GET",
			Controller:  u.GetUser,
			Description: "This endpint gets the user details",
			Middleware: []gin.HandlerFunc{
				middleware.Auth(u.FE.TokenService),
			},
		},
	})
}
