package config

import "os"

type Config struct {
	APP_PORT   string
	JWT_KEY    string
	JWT_EXPIRY string

	POSTGRES_HOST                    string
	POSTGRES_PORT                    string
	POSTGRES_USER                    string
	POSTGRES_PASS                    string
	POSTGRES_DB_NAME                 string
	POSTGRES_SSLMODE                 string
	POSTGRES_CONNECTION_MAX_LIFETIME string
	POSTGRES_MAX_OPEN_CONNECTION     string
	POSTGRES_MAX_IDLE_CONNECTION     string
}

func GenerateConfig() *Config {
	var config Config

	config.APP_PORT = os.Getenv(APP_PORT)
	config.JWT_EXPIRY = os.Getenv(JWT_EXPIRY)
	config.JWT_KEY = os.Getenv(JWT_KEY)
	config.POSTGRES_HOST = os.Getenv(POSTGRES_HOST)
	config.POSTGRES_PORT = os.Getenv(POSTGRES_PORT)
	config.POSTGRES_USER = os.Getenv(POSTGRES_USER)
	config.POSTGRES_PASS = os.Getenv(POSTGRES_PASS)
	config.POSTGRES_DB_NAME = os.Getenv(POSTGRES_DB_NAME)
	config.POSTGRES_SSLMODE = os.Getenv(POSTGRES_SSLMODE)
	config.POSTGRES_CONNECTION_MAX_LIFETIME = os.Getenv(POSTGRES_CONNECTION_MAX_LIFETIME)
	config.POSTGRES_MAX_OPEN_CONNECTION = os.Getenv(POSTGRES_MAX_OPEN_CONNECTION)
	config.POSTGRES_MAX_IDLE_CONNECTION = os.Getenv(POSTGRES_MAX_IDLE_CONNECTION)

	return &config
}
