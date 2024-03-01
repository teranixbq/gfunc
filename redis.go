package gfunc

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type rdb struct {
	r *redis.Client
}

type rdbInterface interface {
	Set(key string, value interface{}) error
	SetEx(key string, value interface{}, exp time.Duration) error
}

func NewRedis(r *redis.Client) rdbInterface {
	return &rdb{
		r: r,
	}
}

func (r *rdb) Set(key string, value interface{}) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = r.r.Set(ctx, key, v, 0).Err()
	if err != nil {
		return err
	}

	return err
}

func (r *rdb) SetEx(key string, value interface{}, exp time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = r.r.SetEx(ctx, key, v, exp).Err()
	if err != nil {
		return err
	}

	return err
}

func (r *rdb) Get(key string, data interface{}) (interface{}, error) {
	ctx := context.Background()
	v, err := r.r.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	
	err = json.Unmarshal([]byte(v), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
