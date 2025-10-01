package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"tracker/internal/config"
	"tracker/internal/dto"
	"errors"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct{
	Redis *redis.Client
}

func NewRedisRepo(c *config.Config) *RedisRepo{
	dsn := fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)

	rdb := redis.NewClient(&redis.Options{
		Addr: dsn,
		Password: c.RedisPass,
		DB: 0,
	})

	return &RedisRepo{Redis: rdb}
}

func (r *RedisRepo) SaveUser(uuid string, userSession *dto.UserSession) error {
	ctx := context.Background()

	data, err := json.Marshal(userSession)

	if err != nil{
		return err
	}

	return r.Redis.Set(ctx, uuid, data, 24 *time.Hour).Err()
}

func (r *RedisRepo) DeleteUser(uuid string) error {
	ctx := context.Background()

	return r.Redis.Del(ctx, uuid).Err()
}

func (r *RedisRepo) GetUser(uuid string) (*dto.UserSession, error) {
	ctx := context.Background()

	res, err := r.Redis.Get(ctx, uuid).Result()

	if err == redis.Nil{
		return nil, errors.New("not record found on redis")
	}

	if err != nil{
		return nil, err
	}

	userSession := &dto.UserSession{}

	err = json.Unmarshal([]byte(res), &userSession)

	if err != nil{
		return nil, err
	}
	fmt.Println(userSession)
	return userSession, nil
}