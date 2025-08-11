package user

import (
	"github.com/gin-gonic/gin"
	DBAdapter "github.com/pawannn/famly/adapter/database/postgres"
	metadbAdapter "github.com/pawannn/famly/adapter/metadb/redis"
	smsAdapter "github.com/pawannn/famly/adapter/sms/twillo"
	middleware "github.com/pawannn/famly/middleware"
	httpEngine "github.com/pawannn/famly/pkg/httpEnginer"
	databasePort "github.com/pawannn/famly/port/database"
	metadbPort "github.com/pawannn/famly/port/metadb"
	"github.com/pawannn/famly/port/sms"
)

type User struct {
	FE            httpEngine.FamlyEngine
	UserRepo      *databasePort.UserDBport
	UserCacheRepo metadbPort.UserCachePort
	UserSmsRepo   sms.UserSmsPort
}

func InitUserRepo(fE httpEngine.FamlyEngine) *User {
	// Initialize user Cache service
	userCacheService := metadbAdapter.InitUserCacheRepo(fE.MetaDB)
	userCacheRepo := metadbPort.InitUserCachePort(userCacheService)

	// Initialize user DB service
	userDBService := DBAdapter.NewUserDBRepository(fE.DB)
	userDBRepo := databasePort.InitUserDBPort(userDBService)

	// Initialize sms service
	UsersmsService := smsAdapter.InitUserSmsRepo(fE.Sms)
	userSmsRepo := sms.InitUserSmsPort(UsersmsService)

	return &User{
		FE:            fE,
		UserRepo:      userDBRepo,
		UserCacheRepo: userCacheRepo,
		UserSmsRepo:   userSmsRepo,
	}
}

func (u *User) InitUserRoutes() {
	u.FE.AddRoute([]httpEngine.FamlyRoute{
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
