package metadbdomain

import "time"

type MetaDBRepository interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiry time.Duration) error
	Delete(key string) error
}
