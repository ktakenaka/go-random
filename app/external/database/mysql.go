package database

import (
	"log"

	"github.com/jinzhu/gorm"
	// http://gorm.io/docs/connecting_to_the_database.html
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/ktakenaka/go-random/app/config"
)

func MySQLConnection() *gorm.DB {
	db, err := gorm.Open("mysql", config.DBACCESS)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
