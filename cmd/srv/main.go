package main

import (
	"github.com/ktakenaka/go-random/app/interface/api"
)

func main() {
	router := api.Handler()
	router.Run()
}
