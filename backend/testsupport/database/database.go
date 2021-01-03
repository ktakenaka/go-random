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
	// TODO: DBからSQLでテーブルを取得するように変更
	tables := []string{"users", "samples"}
	for _, t := range tables {
		if err := d.Session().Exec("DELETE FROM " + t).Error; err != nil {
			return err
		}
	}
	return nil
}
