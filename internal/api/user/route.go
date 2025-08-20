package user

import (
	"github.com/gin-gonic/gin"
	DBAdapter "github.com/pawannn/famly/internal/adapter/database/postgres"
	metadbAdapter "github.com/pawannn/famly/internal/adapter/metadb/redis"
	smsAdapter "github.com/pawannn/famly/internal/adapter/sms/twillo"
	datastoredomain "github.com/pawannn/famly/internal/core/domain/datastore"
	metadbdomain "github.com/pawannn/famly/internal/core/domain/metadb"
	middleware "github.com/pawannn/famly/internal/middleware"
	httpEngine "github.com/pawannn/famly/internal/pkg/httpEnginer"
)

type User struct {
	FE            httpEngine.FamlyEngine
	UserRepo      datastoredomain.UserDBRepo
	UserCacheRepo metadbdomain.UserMetaDBRepo
	UserSmsRepo   smsAdapter.UserSmsRepo
}

func InitUserRepo(fE httpEngine.FamlyEngine) *User {
	// Initialize user DB service
	userDBService := DBAdapter.NewUserDBRepository(fE.DB)
	// Initialize user Cache service
	userCacheService := metadbAdapter.InitUserMetaDbRepo(*fE.MetaDB)
	// Initialize sms service
	UsersmsService := smsAdapter.InitUserSmsRepo(fE.Sms)

	return &User{
		FE:            fE,
		UserRepo:      userDBService,
		UserCacheRepo: userCacheService,
		UserSmsRepo:   UsersmsService,
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
				middleware.Auth(u.FE.Auth),
			},
		},
		{
			Route:       "/user",
			Method:      "GET",
			Controller:  u.GetUser,
			Description: "This endpint gets the user details",
			Middleware: []gin.HandlerFunc{
				middleware.Auth(u.FE.Auth),
			},
		},
	})
}
