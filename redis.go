package gfunc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	r *redis.Client
}

func NewRedis(r *redis.Client) *Redis {
	return &Redis{
		r: r,
	}
}

func (r *Redis) SetJSON(key string, value interface{}) error {
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

func (r *Redis) SetExJSON(key string, value interface{}, exp time.Duration) error {
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

func (r *Redis) GetJSON(key string, data interface{}) error {
	ctx := context.Background()
	v, err := r.r.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if reflect.TypeOf(data).Kind() == reflect.Struct {
		if reflect.TypeOf(data).Kind() != reflect.Ptr {
			err := ErrMsg("data must be of type pointer if struct")
			log.SetFlags(0)
			log.Println(err)
		}
	}

	err = json.Unmarshal([]byte(v), data)
	if err != nil {
		return err
	}

	return nil
}

func ErrMsg(err string) string {
	const redColor = "\033[31m"
	const resetColor = "\033[0m"

	response := fmt.Sprintf("%serror%s: %s", redColor, resetColor, err)
	return response
}
