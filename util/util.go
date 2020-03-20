package util

import "os"

func GetHost() string {
	if host := os.Getenv("DATABASE_URL"); host != "" {
		return host
	}
	return "host=db user=prefectures password=root dbname=prefectures port=5432 sslmode=disable"
}
