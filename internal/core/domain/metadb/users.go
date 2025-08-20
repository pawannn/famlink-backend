package metadbdomain

import datastoredomain "github.com/pawannn/famly/internal/core/domain/datastore"

type UserMetaDBRepo interface {
	SetUser(userid string, userDetails datastoredomain.UserSchema) error
	GetUser(userid string) (*datastoredomain.UserSchema, error)
	DeleteUser(userid string) error
}
