package redis_repositories

import (
	"context"
	"github.com/redis/go-redis/v9"
	"src/internal/models"
	"src/internal/repositories"
	"time"
)

var ctx = context.Background()

type PasswordRecordRedisRepository struct {
	Client        *redis.Client
	ExpireSeconds time.Duration
}

func NewPasswordRecordRedisRepository(fields *RedisRepositoryFields) repositories.PasswordRecordRepository {
	return &PasswordRecordRedisRepository{Client: fields.Client, ExpireSeconds: fields.ExpireSeconds}
}

func (r *PasswordRecordRedisRepository) Set(record *models.PasswordRecord) error {
	key := CompositeKey(record.Username, record.Service)
	err := r.Client.Set(ctx, key, record.Password, r.ExpireSeconds).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *PasswordRecordRedisRepository) Get(username string, service string) (*models.PasswordRecord, error) {
	key := CompositeKey(username, service)
	res, err := r.Client.Get(ctx, key).Result()
	if err != nil && err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	record := &models.PasswordRecord{
		Username: username,
		Service:  service,
		Password: res,
	}
	return record, nil
}

func (r *PasswordRecordRedisRepository) Delete(username string, service string) error {
	key := CompositeKey(username, service)
	_, err := r.Client.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}
