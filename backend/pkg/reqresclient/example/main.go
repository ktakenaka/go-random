package main

import (
	"context"
	"fmt"

	client "github.com/ktakenaka/go-random/backend/pkg/reqresclient"
)

func main() {
	conf := client.NewReqResConfig("https://reqres.in/api")
	user, err := conf.GetUser(context.Background(), 10)

	fmt.Println(err)
	fmt.Println(user)
}
