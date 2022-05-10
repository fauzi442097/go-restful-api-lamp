package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func Env(key, defaultValue string) string {
	if value, isExists := os.LookupEnv(key); isExists {
		return value
	}

	return defaultValue
}
