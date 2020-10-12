package database

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once sync.Once
	db   *gorm.DB
)

// InitMySQLConnection init db connection
func InitMySQLConnection(confStr string) (err error) {
	once.Do(func() {
		db, err = gorm.Open(
			mysql.Open(confStr),
			&gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			},
		)
	})
	if err != nil {
		return err
	}
	return nil
}

// MySQLConnection returns db
func MySQLConnection() *gorm.DB {
	return db
}
