package main

import (
	"log"
	"os"

	"github.com/ktakenaka/go-random/app/config"
	"github.com/ktakenaka/go-random/app/external/database"
	"github.com/ktakenaka/go-random/app/external/framework"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "go-random",
		Usage: "try everything",
		Action: func(c *cli.Context) error {
			if err := database.InitMySQLConnection(config.GenDBAceessStr()); err != nil {
				return err
			}

			router := framework.Handler()
			if err := router.Run(":8080"); err != nil {
				return err
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
