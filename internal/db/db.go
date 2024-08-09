package db

import (
	"context"
	"fmt"

	"github.com/notnmeyer/dngl/internal/envhelper"

	"github.com/redis/go-redis/v9"
)

var keyName = "notes"
var ctx = context.TODO()

type DB struct {
	client *redis.Client
}

func New() *DB {
	return &DB{
		client: redis.NewClient(&redis.Options{
			Addr:     envhelper.New().REDIS_DB_URL, // "redis:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (db *DB) Save(field, value string) error {
	_, err := db.client.HSet(ctx, keyName, field, value).Result()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Get(field string) (*string, error) {
	result, err := db.client.HGet(ctx, keyName, field).Result()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (db *DB) Delete(field string) error {
	delCount, err := db.client.HDel(ctx, keyName, field).Result()
	if err != nil {
		return err
	}

	if delCount == 0 {
		return fmt.Errorf("no fields matching '%s' found", field)
	}

	return nil
}

func (db *DB) GetAll() (map[string]string, error) {
	result, err := db.client.HGetAll(ctx, keyName).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}
