package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	APP_PORT   string
	JWT_KEY    string
	JWT_EXPIRY string

	CLIENT_HOST    string
	CLIENT_PORT    string
	CLIENT_TIMEOUT time.Duration
}

func GenerateConfig() *Config {
	var config Config

	clientTimeoutStr := os.Getenv(CLIENT_TIMEOUT)
	clientTimeout, _ := strconv.Atoi(clientTimeoutStr)

	config.APP_PORT = os.Getenv(APP_PORT)
	config.JWT_EXPIRY = os.Getenv(JWT_EXPIRY)
	config.JWT_KEY = os.Getenv(JWT_KEY)
	config.CLIENT_HOST = os.Getenv(CLIENT_HOST)
	config.CLIENT_PORT = os.Getenv(CLIENT_PORT)
	config.CLIENT_TIMEOUT = time.Second * time.Duration(clientTimeout)

	return &config
}
