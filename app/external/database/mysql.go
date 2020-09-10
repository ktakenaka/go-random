package database

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

// pass config.DBACCESS to this function in main()
func InitMySQLConnection(confStr string) (err error) {
	once.Do(func() {
		db, err = gorm.Open(mysql.Open(confStr), &gorm.Config{})
	})
	if err != nil {
		return err
	}
	return nil
}

func MySQLConnection() *gorm.DB {
	return db
}
