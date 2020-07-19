package config

import (
	"log"

	"github.com/jinzhu/gorm"
	// http://gorm.io/docs/connecting_to_the_database.html
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBConnection() *gorm.DB {
	// TODO: use environment variable to access DB
	db, err := gorm.Open("mysql", "random:random@tcp(db:3306)/go-random?parseTime=true")

	// TODO: close db connection
	// defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	return db
}
