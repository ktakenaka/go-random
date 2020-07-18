package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBConnection() *gorm.DB {
	// TODO: use environment variable to access DB
	db, err := gorm.Open("mysql", "random:random@tcp(db:3306)/go-random")

	// TODO: close db connection
	//defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	return db
}
