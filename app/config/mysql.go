package config

import (
	"fmt"
	"os"
)

var USER = os.Getenv("MYSQL_USER")
var PASSWORD = os.Getenv("MYSQL_PASSWORD")
var DBHOST = os.Getenv("DBHOST")

func GenDBAceessStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/go-random?parseTime=true&charset=utf8mb4",
		USER, PASSWORD, DBHOST)
}
