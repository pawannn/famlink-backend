package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	metadb "github.com/pawannn/famlink/core/services/metaDB"
	appconfig "github.com/pawannn/famlink/pkg/appConfig"
	"github.com/redis/go-redis/v9"
)

type CacheRepo struct {
	Client *redis.Client
}

func InitCacheRepo(c appconfig.Config) metadb.MetaDBService {
	addr := fmt.Sprintf("%s:%d", c.MetaDB_Host, c.MetaDB_Port)
	rds := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: c.DB_pass,
		DB:       c.MetaDB_DB,
	})
	cR := CacheRepo{
		Client: rds,
	}
	return cR
}

func (cR CacheRepo) Set(key string, value interface{}) error {
	ctx := context.Background()
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	cmd := cR.Client.Set(ctx, key, data, time.Minute*10)
	return cmd.Err()
}

func (cR CacheRepo) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := cR.Client.Get(ctx, key).Result()
	if err != redis.Nil {
		return "", err
	}
	return val, nil
}

func (cR CacheRepo) Delete(key string) error {
	ctx := context.Background()
	return cR.Client.Del(ctx, key).Err()
}
