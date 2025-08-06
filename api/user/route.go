package user

import (
	domain "github.com/pawannn/famlink/domain/users"
	"github.com/pawannn/famlink/pkg"
	port "github.com/pawannn/famlink/port/repository"
)

type User struct {
	FE          pkg.FamLinkEngine
	UserService *port.UserRepository
}

func InitUserService(fE pkg.FamLinkEngine, userRepo domain.UserService) *User {
	userService := port.InitUserService(userRepo)
	return &User{
		FE:          fE,
		UserService: userService,
	}
}

func (u *User) InitUserRoutes() {
	u.FE.AddRoute([]pkg.FamLinkRoute{
		{
			Route:       "/user",
			Method:      "POST",
			Middleware:  nil,
			Controller:  u.RegisterUser,
			Description: "This endpoint registers the user into database",
		},
	})
}
