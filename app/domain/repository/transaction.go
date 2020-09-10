package repository

import (
	"gorm.io/gorm"
)

type TransactionManager interface {
	Begin()
	Commit() error
	Rollback()
	GetTx() *gorm.DB
}
