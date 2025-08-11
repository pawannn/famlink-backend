package metadb

import (
	"encoding/json"
	"time"

	domain "github.com/pawannn/famly/core/domain/users"
	metadb "github.com/pawannn/famly/core/services/metaDB"
	"github.com/pawannn/famly/pkg/constants"
)

type UserMetaDBRepo struct {
	rds metadb.MetaDBService
}

func InitUserCacheRepo(rds metadb.MetaDBService) metadb.UserCacheService {
	return UserMetaDBRepo{
		rds: rds,
	}
}

func (uC UserMetaDBRepo) SaveUser(userID string, userDetails domain.UserSchema) error {
	key := userID + constants.USER_METADB_KEY_SUFFIX
	return uC.rds.Set(key, userDetails, time.Minute*10)
}

func (uC UserMetaDBRepo) GetUser(userID string) (*domain.UserSchema, error) {
	key := userID + constants.USER_METADB_KEY_SUFFIX
	data, err := uC.rds.Get(key)
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
	key := userID + constants.USER_METADB_KEY_SUFFIX
	return uC.rds.Delete(key)
}

func (uC UserMetaDBRepo) SetUserOTP(userID string, otp int) error {
	key := userID + constants.USER_METADB_OTP_SUFFIX
	return uC.rds.Set(key, otp, time.Minute*10)
}

func (uC UserMetaDBRepo) GetUserOTP(userID string) (int, error) {
	key := userID + constants.USER_METADB_OTP_SUFFIX
	data, err := uC.rds.Get(key)
	if err != nil {
		return 0, err
	}
	var otp int
	if err := json.Unmarshal([]byte(data), &otp); err != nil {
		return 0, err
	}
	return otp, nil
}

func (uC UserMetaDBRepo) DeleteUserOTP(userID string) error {
	key := userID + constants.USER_METADB_OTP_SUFFIX
	return uC.rds.Delete(key)
}
