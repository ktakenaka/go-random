package database

import (
	"fmt"
	"strings"
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

	Options            map[string]string
	IsIgnoreForeignKey bool
	LogLevel           string
}

// DB db connection
type DB struct {
	session *gorm.DB
}

// TX transaction
type TX struct {
	session *gorm.DB
}

const (
	conn          = "%s:%s@tcp(%s)/%s?collation=utf8mb4_bin&parseTime=true&charset=utf8mb4"
	connWithoutFK = conn + "&foreign_key_checks=0"
)

const (
	defaultMaxIdleConns    = 5
	defaultMaxOpenConns    = 10
	defaultConnMaxLifetime = 20 * time.Second
)

// New connect to db
func New(cfg *Config) *DB {
	var connStr string
	if cfg.IsIgnoreForeignKey {
		connStr = connWithoutFK
	} else {
		connStr = conn
	}

	if cfg.Options != nil {
		var options []string
		for k, v := range cfg.Options {
			options = append(options, k+"="+v)
		}
		connStr = connStr + "&" + strings.Join(options, "&")
	}
	dst := fmt.Sprintf(connStr, cfg.User, cfg.Password, cfg.Host, cfg.Name)

	return &DB{
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

	if cfg.MaxIdleConns == 0 {
		sqldb.SetMaxIdleConns(defaultMaxIdleConns)
	} else {
		sqldb.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	if cfg.MaxOpenConns == 0 {
		sqldb.SetMaxIdleConns(defaultMaxOpenConns)
	} else {
		sqldb.SetMaxOpenConns(cfg.MaxOpenConns)
	}

	if cfg.ConnMaxLifetime == 0 {
		sqldb.SetConnMaxLifetime(defaultConnMaxLifetime)
	} else {
		sqldb.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}
	return db
}

// Session returns tion
func (d *DB) Session() *gorm.DB {
	return d.session
}

// Begin transaction
func (d *DB) Begin() (*TX, error) {
	tx := d.session.Begin()
	return &TX{session: tx}, tx.Error
}

// Close connections
func (d *DB) Close() error {
	sqlDB, err := d.session.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Session returns connection
func (t *TX) Session() *gorm.DB {
	return t.session
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
