package services

import (
	"time"

	metadbdomain "github.com/pawannn/famly/internal/core/domain/metadb"
)

type MetaDBManager struct {
	repo metadbdomain.MetaDBRepository
}

func NewMetaDBManager(repo metadbdomain.MetaDBRepository) *MetaDBManager {
	return &MetaDBManager{repo: repo}
}

func (mdM *MetaDBManager) Get(key string) (string, error) {
	return mdM.repo.Get(key)
}

func (mdM *MetaDBManager) Set(key string, value interface{}, expiry time.Duration) error {
	return mdM.repo.Set(key, value, expiry)
}

func (mdM *MetaDBManager) Delete(key string) error {
	return mdM.repo.Delete(key)
}
