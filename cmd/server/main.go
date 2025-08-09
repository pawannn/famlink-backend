package main

import (
	"log"

	database "github.com/pawannn/famlink/adapter/database/postgres"
	token "github.com/pawannn/famlink/adapter/token/jwt"
	"github.com/pawannn/famlink/api/user"
	appconfig "github.com/pawannn/famlink/pkg/appConfig"
	httpEngine "github.com/pawannn/famlink/pkg/httpEnginer"
	port "github.com/pawannn/famlink/port/token"
)

func main() {
	var c appconfig.Config
	appconfig.LoadConfig(&c)

	db, err := database.InitDatabase(c)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer db.Close()

	ts := token.InitTokenService(c)
	tokenRepo := port.InitTokenRepo(ts)

	famLinkEngine := httpEngine.InitFamLinkEngine(c, db, *tokenRepo)

	userRepo := database.NewUserRepository(db)
	userRoutes := user.InitUserService(*famLinkEngine, userRepo)
	userRoutes.InitUserRoutes()

	famLinkEngine.StartServer()
}
