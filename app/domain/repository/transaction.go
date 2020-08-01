package repository

import (
	"github.com/jinzhu/gorm"
)

type TransactionManager interface {
	Begin()
	Commit() error
	Rollback()
	GetTx() *gorm.DB
}
