package database

import (
	"sync"

	"github.com/jinzhu/gorm"
	// http://gorm.io/docs/connecting_to_the_database.html
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	once sync.Once
	db   *gorm.DB
)

// pass config.DBACCESS to this function in main()
func InitMySQLConnection(confStr string) (err error) {
	once.Do(func() {
		db, err = gorm.Open("mysql", confStr)
	})
	if err != nil {
		return err
	}
	return nil
}

func MySQLConnection() *gorm.DB {
	return db
}
