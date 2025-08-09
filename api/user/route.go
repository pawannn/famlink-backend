package user

import (
	"github.com/gin-gonic/gin"
	DBAdapter "github.com/pawannn/famlink/adapter/database/postgres"
	metadbAdapter "github.com/pawannn/famlink/adapter/metadb/redis"
	smsAdapter "github.com/pawannn/famlink/adapter/sms/twillo"
	middleware "github.com/pawannn/famlink/middleware"
	httpEngine "github.com/pawannn/famlink/pkg/httpEnginer"
	databasePort "github.com/pawannn/famlink/port/database"
	metadbPort "github.com/pawannn/famlink/port/metadb"
	"github.com/pawannn/famlink/port/sms"
)

type User struct {
	FE            httpEngine.FamLinkEngine
	UserRepo      *databasePort.UserDBport
	UserCacheRepo metadbPort.UserCachePort
	UserSmsRepo   sms.UserSmsPort
}

func InitUserRepo(fE httpEngine.FamLinkEngine) *User {
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
