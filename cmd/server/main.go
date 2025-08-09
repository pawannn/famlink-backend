package main

import (
	"log"

	database "github.com/pawannn/famlink/adapter/database/postgres"
	cache "github.com/pawannn/famlink/adapter/metadb/redis"
	sms "github.com/pawannn/famlink/adapter/sms/twillo"
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
	tokenRepo := port.InitTokenPort(ts)

	// Initialize sms service
	sms := sms.InitTwilloClient(c)

	// Initialize the HTPP engine
	famLinkEngine := httpEngine.InitFamLinkEngine(c, db, *tokenRepo, rds, sms)

	// Initialize the user Repo
	userRoutes := userApi.InitUserRepo(*famLinkEngine)

	// Add User Routes
	userRoutes.InitUserRoutes()

	// Start the server
	famLinkEngine.StartServer()
}
