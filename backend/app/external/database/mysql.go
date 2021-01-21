package database

import (
	"time"

	"github.com/ktakenaka/go-random/backend/pkg/infra/database"
)

var (
	db *database.DB
)

// InitMySQLConnection init db connection
func InitMySQLConnection(user, passord, dbhost string) {
	cfg := database.Config{
		User:            user,
		Password:        passord,
		Host:            dbhost,
		Name:            "go-random",
		MaxIdleConns:    40,
		MaxOpenConns:    200,
		ConnMaxLifetime: 30 * time.Second,
		Options:         map[string]string{"interpolateParams": "true"},

		LogLevel: "silent",
	}

	infradb := database.New(&cfg)
	db = infradb
}

// MySQLConnection returns db
func MySQLConnection() *database.DB {
	return db
}
