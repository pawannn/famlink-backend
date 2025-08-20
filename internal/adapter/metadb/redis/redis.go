package metadb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	metadbdomain "github.com/pawannn/famly/internal/core/domain/metadb"
	appconfig "github.com/pawannn/famly/internal/pkg/appConfig"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

func InitRedisRepo(c appconfig.Config) metadbdomain.MetaDBRepository {
	addr := fmt.Sprintf("%s:%d", c.MetaDB_Host, c.MetaDB_Port)
	rds := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: c.DB_pass,
		DB:       c.MetaDB_DB,
	})
	cR := RedisRepo{
		Client: rds,
	}
	return cR
}

func (cR RedisRepo) Set(key string, value interface{}, expiry time.Duration) error {
	ctx := context.Background()
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	cmd := cR.Client.Set(ctx, key, data, expiry)
	return cmd.Err()
}

func (cR RedisRepo) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := cR.Client.Get(ctx, key).Result()
	if err != redis.Nil {
		return "", err
	}
	return val, nil
}

func (cR RedisRepo) Delete(key string) error {
	ctx := context.Background()
	return cR.Client.Del(ctx, key).Err()
}
