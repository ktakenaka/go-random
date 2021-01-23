package database

import (
	"log"
	"os"

	db "github.com/ktakenaka/go-random/backend/pkg/infra/database"
)

func GetDB() (testDB *db.DB, release func()) {
	cfg := db.Config{
		User:               os.Getenv("MYSQL_USER"),
		Password:           os.Getenv("MYSQL_PASSWORD"),
		Host:               os.Getenv("DBHOST"),
		Name:               "go-random_test1", // TODO: 3並列で実行できるようになおす
		IsIgnoreForeignKey: true,
	}

	testDB = db.New(&cfg)
	release = func() {
		if err := cleanDB(testDB); err != nil {
			log.Panic(err)
		}
		if err := testDB.Close(); err != nil {
			log.Panic(err)
		}
	}
	return
}

func cleanDB(d *db.DB) error {
	// When you add a new table and it's master data, you need to escape it
	rows, err := d.Session().Raw(
		"SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES " +
			"WHERE TABLE_SCHEMA='go-random' AND TABLE_NAME != 'schema_migrations'",
	).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var t string
		if err := rows.Scan(&t); err != nil {
			return err
		}
		// nolint:gosec
		if err := d.Session().Raw("DELETE FROM " + t).Error; err != nil {
			return err
		}
	}
	return nil
}
