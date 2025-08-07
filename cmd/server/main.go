package main

import (
	"log"

	database "github.com/pawannn/famlink/adapter/database/postgres"
	token "github.com/pawannn/famlink/adapter/token/jwt"
	"github.com/pawannn/famlink/api/user"
	"github.com/pawannn/famlink/pkg"
	appconfig "github.com/pawannn/famlink/pkg/appConfig"
)

func main() {
	var c appconfig.Config
	appconfig.LoadConfig(&c)
	ts := token.InitTokenService(c)
	db, err := database.InitDatabase(c)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer db.Close()

	famLinkEngine := pkg.InitFamLinkEngine(c, db, ts)

	userRepo := database.NewUserRepository(db)
	userRoutes := user.InitUserService(*famLinkEngine, userRepo)
	userRoutes.InitUserRoutes()

	famLinkEngine.StartServer()
}
