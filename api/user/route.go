package user

import (
	"github.com/gin-gonic/gin"
	CacheAdapter "github.com/pawannn/famlink/adapter/cache/redis"
	domain "github.com/pawannn/famlink/core/domain/users"
	metadb "github.com/pawannn/famlink/core/services/metaDB"
	"github.com/pawannn/famlink/middleware"
	httpEngine "github.com/pawannn/famlink/pkg/httpEnginer"
	port "github.com/pawannn/famlink/port/database"
)

type User struct {
	FE            httpEngine.FamLinkEngine
	UserRepo      *port.UserRepository
	UserCacheRepo metadb.UserCacheService
}

func InitUserService(fE httpEngine.FamLinkEngine, userService domain.UserService) *User {
	userRepo := port.InitUserService(userService)
	userCacheRepo := CacheAdapter.InitUserCacheRepo(fE.MetaDB)
	return &User{
		FE:            fE,
		UserRepo:      userRepo,
		UserCacheRepo: userCacheRepo,
	}
}

func (u *User) InitUserRoutes() {
	u.FE.AddRoute([]httpEngine.FamLinkRoute{
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
				middleware.Auth(u.FE.Token),
			},
		},
		{
			Route:       "/user",
			Method:      "GET",
			Controller:  u.GetUser,
			Description: "This endpint gets the user details",
			Middleware: []gin.HandlerFunc{
				middleware.Auth(u.FE.Token),
			},
		},
	})
}
