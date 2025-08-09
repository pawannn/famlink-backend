package cache

import (
	"encoding/json"
	"time"

	domain "github.com/pawannn/famlink/core/domain/users"
	metadb "github.com/pawannn/famlink/core/services/metaDB"
)

type UserMetaDBRepo struct {
	rds metadb.MetaDBService
}

func InitUserCacheRepo(rds metadb.MetaDBService) metadb.UserCacheService {
	return UserMetaDBRepo{
		rds: rds,
	}
}

func (uC UserMetaDBRepo) SetUser(userID string, userDetails domain.UserSchema) error {
	return uC.rds.Set(userID, userDetails, time.Minute*10)
}

func (uC UserMetaDBRepo) GetUser(userID string) (*domain.UserSchema, error) {
	data, err := uC.rds.Get(userID)
	if err != nil {
		return nil, err
	}
	var userDetails domain.UserSchema
	if err := json.Unmarshal([]byte(data), &userDetails); err != nil {
		return nil, err
	}

	return &userDetails, nil
}

func (uC UserMetaDBRepo) DeleteUser(userID string) error {
	return uC.rds.Delete(userID)
}
