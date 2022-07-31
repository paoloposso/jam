package config

import "os"

func GetMongoUrlAndDatabase() (url string, database string) {
	return os.Getenv("MONGO_URL"), os.Getenv("MONGO_DATABASE")
}

func GetPort() string {
	return os.Getenv("PORT")
}
