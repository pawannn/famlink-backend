package metadb

import "time"

type MetaDBService interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiry time.Duration) error
	Delete(key string) error
}
