package db

import (
	"os"
)

func GetUrlConnection() string {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USERDB")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return "host=" + host +
		" port=" + port +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" sslmode=disable"
}
