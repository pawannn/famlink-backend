package metadb

import (
	"encoding/json"
	"time"

	datastoredomain "github.com/pawannn/famly/internal/core/domain/datastore"
	metadbdomain "github.com/pawannn/famly/internal/core/domain/metadb"
	"github.com/pawannn/famly/internal/core/services"
	"github.com/pawannn/famly/internal/pkg/constants"
)

type UserMetaDBRepo struct {
	rds services.MetaDBManager
}

func InitUserMetaDbRepo(rds services.MetaDBManager) metadbdomain.UserMetaDBRepo {
	return UserMetaDBRepo{
		rds: rds,
	}
}

func (uC UserMetaDBRepo) SetUser(userID string, userDetails datastoredomain.UserSchema) error {
	key := userID + constants.USER_METADB_KEY_SUFFIX
	return uC.rds.Set(key, userDetails, time.Minute*10)
}

func (uC UserMetaDBRepo) GetUser(userID string) (*datastoredomain.UserSchema, error) {
	key := userID + constants.USER_METADB_KEY_SUFFIX
	data, err := uC.rds.Get(key)
	if err != nil {
		return nil, err
	}
	var userDetails datastoredomain.UserSchema
	if err := json.Unmarshal([]byte(data), &userDetails); err != nil {
		return nil, err
	}

	return &userDetails, nil
}

func (uC UserMetaDBRepo) DeleteUser(userID string) error {
	key := userID + constants.USER_METADB_KEY_SUFFIX
	return uC.rds.Delete(key)
}
