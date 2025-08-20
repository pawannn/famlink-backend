package user

import (
	"github.com/pawannn/famly/internal/application/users"
	httpEngine "github.com/pawannn/famly/internal/pkg/httpEnginer"
)

type User struct {
	FE              httpEngine.FamlyEngine
	UserApplication users.UserApplication
}

func InitUserRepo(fE httpEngine.FamlyEngine) *User {
	uA := users.InitUserApplication(*fE.MetaDB, fE.Sms, fE.DB)

	return &User{
		FE:              fE,
		UserApplication: uA,
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
		// {
		// 	Route:       "/user/phone/verify",
		// 	Method:      "PUT",
		// 	Controller:  u.VerifyPhone,
		// 	Description: "This endpoint verify the phone with the given OTP",
		// 	Middleware:  nil,
		// },
		// {
		// 	Route:       "/user",
		// 	Method:      "PATCH",
		// 	Controller:  u.UpdateUser,
		// 	Description: "This endpoint updates the username of the user",
		// 	Middleware: []gin.HandlerFunc{
		// 		middleware.Auth(u.FE.Auth),
		// 	},
		// },
		// {
		// 	Route:       "/user",
		// 	Method:      "GET",
		// 	Controller:  u.GetUser,
		// 	Description: "This endpint gets the user details",
		// 	Middleware: []gin.HandlerFunc{
		// 		middleware.Auth(u.FE.Auth),
		// 	},
		// },
	})
}
