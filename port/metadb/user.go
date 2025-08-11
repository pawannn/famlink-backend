package metadb

import (
	domain "github.com/pawannn/famly/core/domain/users"
	metadb "github.com/pawannn/famly/core/services/metaDB"
)

type UserCachePort struct {
	Repo metadb.UserCacheService
}

func InitUserCachePort(repo metadb.UserCacheService) UserCachePort {
	ucR := UserCachePort{
		Repo: repo,
	}
	return ucR
}

func (ucR UserCachePort) GetUser(userID string) (*domain.UserSchema, error) {
	return ucR.Repo.GetUser(userID)
}

func (ucR UserCachePort) SaveUser(userID string, userDetails domain.UserSchema) error {
	return ucR.Repo.SaveUser(userID, userDetails)
}

func (ucR UserCachePort) DeleteUser(userID string) error {
	return ucR.Repo.DeleteUser(userID)
}
