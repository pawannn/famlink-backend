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
	// Load Env
	var c appconfig.Config
	appconfig.LoadConfig(&c)

	// Initialize Database
	db, err := database.InitDatabase(c)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer db.Close()

	// Initialize token service
	ts := token.InitTokenService(c)
	tokenRepo := port.InitTokenRepo(ts)

	// Initialize the HTPP engine
	famLinkEngine := httpEngine.InitFamLinkEngine(c, db, *tokenRepo)

	// Initialize the user Repo
	userRepo := database.NewUserRepository(db)
	userRoutes := user.InitUserService(*famLinkEngine, userRepo)

	// Add User Routes
	userRoutes.InitUserRoutes()

	// Start the server
	famLinkEngine.StartServer()
}
