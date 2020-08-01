package main

import (
	"log"

	"github.com/ktakenaka/go-random/app/config"
	"github.com/ktakenaka/go-random/app/external/database"
	"github.com/ktakenaka/go-random/app/external/framework"
)

func main() {
	if err := database.InitMySQLConnection(config.GenDBAceessStr()); err != nil {
		log.Fatal(err)
	}

	router := framework.Handler()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
