package database

import (
	"time"

	"gorm.io/gorm"

	"github.com/ktakenaka/go-random/backend/pkg/infra/database"
)

var (
	db database.DB
)

// InitMySQLConnection init db connection
func InitMySQLConnection(user, passord, dbhost string) (err error) {
	cfg := database.Config{
		User:            user,
		Password:        passord,
		Host:            dbhost,
		Name:            "go-random",
		MaxIdleConns:    40,
		MaxOpenConns:    200,
		ConnMaxLifetime: 30 * time.Second,
	}

	db = database.New(&cfg)
	return nil
}

// MySQLConnection returns db
func MySQLConnection() *gorm.DB {
	// TODO: repositoryでもdatabase.DB使うようにする
	return db.Session()
}
