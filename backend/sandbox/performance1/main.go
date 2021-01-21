package main

import (
	"fmt"
	"time"

	"github.com/ktakenaka/go-random/backend/pkg/infra/database"
)

func main() {
	cfg := database.Config{
		User:     "random",
		Password: "random",
		Host:     "db",
		Name:     "go-random",
		LogLevel: "silent",
	}
	fmt.Println("--- without interpolate ---")
	measure(cfg)

	cfgInterpolate := database.Config{
		User:     "random",
		Password: "random",
		Host:     "db",
		Name:     "go-random",
		Options:  map[string]string{"interpolateParams": "true"},
		LogLevel: "silent",
	}
	fmt.Println("--- with interpolate ---")
	measure(cfgInterpolate)
}

func measure(cfg database.Config) {
	db := database.New(&cfg)
	db.Session().Exec(
		`CREATE TABLE IF NOT EXISTS tmp (
			id INTEGER UNSIGNED NOT NULL,
			name VARCHAR(100)
		)`,
	)
	start := time.Now()
	for i := 1; i < 1000; i++ {
		db.Session().Exec("INSERT INTO tmp VALUES(?, ?)", i, "hoge")
	}
	end := time.Now()
	diff := end.Sub(start)
	fmt.Println(diff)
	db.Session().Exec(
		`DROP TABLE tmp`,
	)
	db.Close()
}
