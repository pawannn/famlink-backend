package main

import (
	"log"

	cache "github.com/pawannn/famlink/adapter/cache/redis"
	database "github.com/pawannn/famlink/adapter/database/postgres"
	token "github.com/pawannn/famlink/adapter/token/jwt"
	userApi "github.com/pawannn/famlink/api/user"
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

	// Intiialize MetaDB
	rds := cache.InitCacheRepo(c)

	// Initialize token service
	ts := token.InitTokenService(c)
	tokenRepo := port.InitTokenRepo(ts)

	// Initialize the HTPP engine
	famLinkEngine := httpEngine.InitFamLinkEngine(c, db, *tokenRepo, rds)

	// Initialize the user Repo
	userRepo := database.NewUserRepository(db)
	userRoutes := userApi.InitUserService(*famLinkEngine, userRepo)

	// Add User Routes
	userRoutes.InitUserRoutes()

	// Start the server
	famLinkEngine.StartServer()
}
