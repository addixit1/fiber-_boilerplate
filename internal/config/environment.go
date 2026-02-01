package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env       string
	Port      string
	MongoURI  string
	JWTSecret string
	RedisURI  string
	MongoDbName string
	DebugStatus string
}

var Config AppConfig

func LoadEnv() {
	_ = godotenv.Load()

	Config = AppConfig{
		Env:       getEnv("ENV", "development"),
		Port:      getEnv("PORT", "3010"),
		MongoURI:  getEnv("MONGO_URI", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
		RedisURI:  getEnv("REDIS_URL", "localhost:6379"),
		MongoDbName: getEnv("MONGO_DB_NAME", ""),
		DebugStatus: getEnv("MONGO_DEBUG", "false"),
	}

	if Config.MongoURI == "" {
		log.Fatal("MONGO_URI is required")
	}
	if Config.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
