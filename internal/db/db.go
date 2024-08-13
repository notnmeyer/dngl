package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/notnmeyer/dngl/internal/envhelper"

	"github.com/redis/go-redis/v9"
)

var defaultExpiry = time.Minute * 15
var ctx = context.TODO()
var keyPrefix = "notes:"

type DB struct {
	client *redis.Client
}

func key(s string) string {
	if strings.HasPrefix(s, keyPrefix) {
		return s
	}
	return keyPrefix + s
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

func (db *DB) Save(id, value string) error {
	_, err := db.client.Set(ctx, key(id), value, defaultExpiry).Result()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Get(id string) (*string, error) {
	fmt.Println("called Get() with " + id)
	result, err := db.client.Get(ctx, key(id)).Result()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (db *DB) Delete(id string) error {
	delCount, err := db.client.Del(ctx, key(id)).Result()
	if err != nil {
		return err
	}

	if delCount == 0 {
		return fmt.Errorf("no record matching '%s' found", key(id))
	}

	return nil
}

func (db *DB) GetAll() ([]string, error) {
	// glob := keyPrefix + "*"
	keys, _, err := db.client.Scan(ctx, 0, "notes:*", 100).Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}
