package database

import (
	"time"

	"github.com/ktakenaka/go-random/backend/pkg/infra/database"
)

var (
	// TODO: pointerじゃないとだめか検討
	// 中の値がpointerなので構造体自体は値で良いかも
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
	}

	infradb := database.New(&cfg)
	db = &infradb
}

// MySQLConnection returns db
func MySQLConnection() *database.DB {
	return db
}
