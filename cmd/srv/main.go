package main

import (
	"github.com/ktakenaka/go-random/app/external/framework"

	"log"
)

func main() {
	router := framework.Handler()
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
