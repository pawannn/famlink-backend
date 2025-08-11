package main

import (
	"log"

	database "github.com/pawannn/famly/adapter/database/postgres"
	cache "github.com/pawannn/famly/adapter/metadb/redis"
	sms "github.com/pawannn/famly/adapter/sms/twillo"
	token "github.com/pawannn/famly/adapter/token/jwt"
	userApi "github.com/pawannn/famly/api/user"
	appconfig "github.com/pawannn/famly/pkg/appConfig"
	httpEngine "github.com/pawannn/famly/pkg/httpEnginer"
	port "github.com/pawannn/famly/port/token"
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
	famlyEngine := httpEngine.InitfamlyEngine(c, db, *tokenRepo, rds, sms)

	// Initialize the user Repo
	userRoutes := userApi.InitUserRepo(*famlyEngine)

	// Add User Routes
	userRoutes.InitUserRoutes()

	// Start the server
	famlyEngine.StartServer()
}
