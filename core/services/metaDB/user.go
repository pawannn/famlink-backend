package metadb

import domain "github.com/pawannn/famly/core/domain/users"

type UserCacheService interface {
	GetUser(userID string) (*domain.UserSchema, error)
	SaveUser(userID string, userDetails domain.UserSchema) error
	DeleteUser(userID string) error
	SetUserOTP(userID string, otp int) error
	GetUserOTP(userID string) (int, error)
	DeleteUserOTP(userID string) error
}
