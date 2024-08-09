package envhelper

import (
	"fmt"
	"os"

	// will automatically load .env if present
	_ "github.com/joho/godotenv/autoload"
)

var requiredEnv = []string{
	"REDIS_DB_URL",
	"DNGL_API_HOST",
	"DNGL_API_PORT",
	"DNGL_TOKEN",
}

type Env struct {
	REDIS_DB_URL  string
	DNGL_API_PORT string
	DNGL_API_URL  string
	DNGL_TOKEN    string
}

func init() {
	for _, env := range requiredEnv {
		_, exists := os.LookupEnv(env)
		if !exists {
			panic(fmt.Sprintf("required env var '%s' not found", env))
		}
	}
}

func New() *Env {
	return &Env{
		REDIS_DB_URL:  os.Getenv("REDIS_DB_URL"),
		DNGL_API_PORT: os.Getenv("DNGL_API_PORT"),
		DNGL_API_URL:  fmt.Sprintf("http://%s:%s", os.Getenv("DNGL_API_HOST"), os.Getenv("DNGL_API_PORT")),
		DNGL_TOKEN:    os.Getenv("DNGL_TOKEN"),
	}
}
