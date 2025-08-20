package users

import (
	"database/sql"

	database "github.com/pawannn/famly/internal/adapter/database/postgres"
	metadb "github.com/pawannn/famly/internal/adapter/metadb/redis"
	datastoredomain "github.com/pawannn/famly/internal/core/domain/datastore"
	metadbdomain "github.com/pawannn/famly/internal/core/domain/metadb"
	"github.com/pawannn/famly/internal/core/services"
)

type UserApplication struct {
	UserDBRepo    datastoredomain.UserDBRepo
	UserCacheRepo metadbdomain.UserMetaDBRepo
	SmsRepo       services.SmsManager
}

func InitUserApplication(metaDbManager services.MetaDBManager, smsManager services.SmsManager, db *sql.DB) UserApplication {
	userDBRepo := database.NewUserDBRepository(db)
	userCacheRepo := metadb.InitUserMetaDbRepo(metaDbManager)
	return UserApplication{
		UserDBRepo:    userDBRepo,
		UserCacheRepo: userCacheRepo,
		SmsRepo:       smsManager,
	}
}
