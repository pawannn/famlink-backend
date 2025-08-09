package metadb

import (
	domain "github.com/pawannn/famlink/core/domain/users"
	metadb "github.com/pawannn/famlink/core/services/metaDB"
)

type UserCacheRepo struct {
	Repo metadb.UserCacheService
}

func InitUserCacheRepo(repo metadb.UserCacheService) UserCacheRepo {
	ucR := UserCacheRepo{
		Repo: repo,
	}
	return ucR
}

func (ucR UserCacheRepo) GetUser(userID string) (*domain.UserSchema, error) {
	return ucR.Repo.GetUser(userID)
}

func (ucR UserCacheRepo) SetUser(userID string, userDetails domain.UserSchema) error {
	return ucR.Repo.SetUser(userID, userDetails)
}

func (ucR UserCacheRepo) DeleteUser(userID string) error {
	return ucR.Repo.DeleteUser(userID)
}
