package redis_repositories

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"src/config"
	"src/internal/repositories"
	"time"
)

type RedisRepositoryFields struct {
	Client        *redis.Client
	ExpireSeconds time.Duration
}

func CreateRedisRepositoryFields(fileName, filePath string) (*RedisRepositoryFields, error) {
	fields := new(RedisRepositoryFields)
	var cfg config.Config
	err := cfg.ParseConfig(fileName, filePath)
	if err != nil {
		return nil, err
	}

	fields.Client, err = cfg.Redis.Init()
	if err != nil {
		return nil, err
	}
	fields.ExpireSeconds = time.Second * time.Duration(cfg.Redis.ExpireSeconds)
	return fields, nil
}

func CreatePasswordRecordRedisRepository(fields *RedisRepositoryFields) repositories.PasswordRecord {
	return NewPasswordRecordRedisRepository(fields)
}

func CompositeKey(s1, s2 string) string {
	return fmt.Sprintf("{%s}%s", s1, s2)
}
