package metadb

import domain "github.com/pawannn/famlink/core/domain/users"

type UserCacheService interface {
	GetUser(userID string) (*domain.UserSchema, error)
	SetUser(userID string, userDetails domain.UserSchema) error
	DeleteUser(userID string) error
}
