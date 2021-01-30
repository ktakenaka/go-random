package main

import (
	"log"
	"os"

	"github.com/ktakenaka/go-random/backend/app/config"
	"github.com/ktakenaka/go-random/backend/app/external/database"
	"github.com/ktakenaka/go-random/backend/app/external/framework"
	"github.com/ktakenaka/go-random/backend/app/interface/api/middleware"
	"github.com/ktakenaka/go-random/backend/pkg/jwtutil"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "go-random",
		Usage: "try everything",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "env",
				EnvVars: []string{"ENV"},
				Value:   "development",
			},
			&cli.StringFlag{
				Name:    "mysqluser",
				EnvVars: []string{"MYSQL_USER"},
			},
			&cli.StringFlag{
				Name:    "mysqlpassword",
				EnvVars: []string{"MYSQL_PASSWORD"},
			},
			&cli.StringFlag{
				Name:    "dbhost",
				EnvVars: []string{"DBHOST"},
			},
			&cli.StringFlag{
				Name:    "google_client_id",
				EnvVars: []string{"GOOGLE_CLIENT_ID"},
			},
			&cli.StringFlag{
				Name:    "google_client_secret",
				EnvVars: []string{"GOOGLE_CLIENT_SECRET"},
			},
			&cli.StringFlag{
				Name:    "google_redirect_url",
				EnvVars: []string{"GOOGLE_REDIRECT_URL"},
				Value:   "http://localhost:3000/auth/google/callback",
			},
			&cli.StringFlag{
				Name:    "jwt_secret",
				EnvVars: []string{"JWT_SECRET"},
				Value:   "N74Yq7tIfzO9muKLUhJuWXBZNhd9HXD7",
			},
			&cli.BoolFlag{
				Name: "pprof_enabled",
				EnvVars: []string{"PPROF_ENABLED"},
				Value: false,
			},
		},
		Action: func(c *cli.Context) error {
			database.InitMySQLConnection(
				c.String("mysqluser"),
				c.String("mysqlpassword"),
				c.String("dbhost"),
			)

			config.InitGoogleOIDCCnf(
				c.String("google_redirect_url"),
				c.String("google_client_id"),
				c.String("google_client_secret"),
			)
			jwtutil.InitJWTSecret(
				c.String("jwt_secret"),
			)
			middleware.InitJWTCookieOpt(c.String("env"))

			router := framework.Handler(c.Bool("pprof_enabled"))
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
