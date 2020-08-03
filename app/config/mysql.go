package config

import (
	"fmt"
)

func GenDBAceessStr(user, password, host string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/go-random?parseTime=true&charset=utf8mb4",
		user, password, host)
}
