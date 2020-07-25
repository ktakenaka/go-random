package config

import (
	"os"
)

var USER = os.Getenv("MYSQL_USER")
var PASSWORD = os.Getenv("MYSQL_PASSWORD")
var DBHOST = os.Getenv("DBHOST")
var DBACCESS = USER + ":" + PASSWORD + "@tcp(" + DBHOST + ")/go-random?parseTime=true"
