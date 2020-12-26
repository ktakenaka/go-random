package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config connection information
type Config struct {
	User            string
	Password        string
	Host            string
	Name            string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration

	LogLevel string
}

// DB db connection
type DB struct {
	session *gorm.DB
}

// TX transaction
type TX struct {
	session *gorm.DB
}

// New connect to db
func New(cfg *Config) DB {
	dst := fmt.Sprintf("%s:%s@tcp(%s)/%s?collation=utf8mb4_bin&parseTime=true&charset=utf8mb4", cfg.User, cfg.Password, cfg.Host, cfg.Name)
	return DB{
		session: connect(dst, cfg),
	}
}

func connect(dst string, cfg *Config) *gorm.DB {
	var loglevel logger.LogLevel
	switch cfg.LogLevel {
	case "silent":
		loglevel = logger.Silent
	case "error":
		loglevel = logger.Error
	case "warn":
		loglevel = logger.Warn
	default:
		loglevel = logger.Info
	}

	db, err := gorm.Open(
		mysql.Open(dst),
		&gorm.Config{
			Logger: logger.Default.LogMode(loglevel),
		},
	)
	if err != nil {
		panic(err)
	}

	sqldb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqldb.SetMaxIdleConns(cfg.MaxIdleConns)
	sqldb.SetMaxOpenConns(cfg.MaxOpenConns)
	sqldb.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	return db
}

// Session returns connection
func (d *DB) Session() *gorm.DB {
	return d.session
}

// Session returns connection
func (t *TX) Session() *gorm.DB {
	return t.session
}

// Begin transaction
func (d *DB) Begin() *gorm.DB {
	tx := d.session.Begin()
	return tx
}

// Commit change
func (t *TX) Commit() error {
	t.session.Commit()
	return t.session.Error
}

// Rollback rollback
func (t *TX) Rollback() error {
	t.session.Rollback()
	return t.session.Error
}
