package config

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DataBaseHost string
	DataBasePort string
	DataBaseUser string
	DataBasePass string
	DataBaseName string
	RedisHost	string
	RedisPort	string
	RedisPass 	string
	GinAddr	string
}

func NewConfig() *Config {
	err := godotenv.Load()

	if err != nil{
		logrus.Fatal("Failed to open .env")
	}

	return &Config{
		DataBaseHost: os.Getenv("DB_HOST"),
		DataBasePort: os.Getenv("DB_PORT"),
		DataBaseUser: os.Getenv("DB_USER"),
		DataBasePass: os.Getenv("DB_PASS"),
		DataBaseName: os.Getenv("DB_NAME"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		RedisPass: os.Getenv("REDIS_PASS"),
		GinAddr: os.Getenv("GIN_ADDR"),
	}
}