package main

import (
	"github.com/ktakenaka/go-random/app/interface/api"

	"log"
)

func main() {
	router := api.Handler()
	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
