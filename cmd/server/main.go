package main

import (
	"log"

	database "github.com/pawannn/famly/internal/adapter/database/postgres"
	metadb "github.com/pawannn/famly/internal/adapter/metadb/redis"
	sms "github.com/pawannn/famly/internal/adapter/sms/twillo"
	auth "github.com/pawannn/famly/internal/adapter/token/jwt"
	"github.com/pawannn/famly/internal/api/user"
	"github.com/pawannn/famly/internal/core/services"
	appconfig "github.com/pawannn/famly/internal/pkg/appConfig"
	httpEngine "github.com/pawannn/famly/internal/pkg/httpEnginer"
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
	rds := metadb.InitRedisRepo(c)
	metadbManager := services.NewMetaDBManager(rds)

	// Initialize token service
	jtS := auth.InitJwtService(c)
	AuthManager := services.InitAuthManager(jtS)

	// Initialize sms service
	tS := sms.InitTwilloClient(c)
	smsManager := services.InitSmsManager(tS)

	// Initialize the HTPP engine
	famlyEngine := httpEngine.InitfamlyEngine(c, db, *AuthManager, metadbManager, smsManager)

	// Initialize the user Repo
	userRoutes := user.InitUserRepo(*famlyEngine)

	// Add User Routes
	userRoutes.InitUserRoutes()

	// Start the server
	famlyEngine.StartServer()
}
