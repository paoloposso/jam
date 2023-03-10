package config

import "os"

func GetPort() string {
	return os.Getenv("PORT")
}

func GetGinMode() string {
	return os.Getenv("GIN_MODE")
}
