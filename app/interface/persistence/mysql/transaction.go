package mysql

import (
	"github.com/jinzhu/gorm"
)

type TransactionManager struct {
	DB *gorm.DB
}

func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{DB: db}
}

func (txm *TransactionManager) Begin() {
	txm.DB = txm.DB.Begin()
}

func (txm *TransactionManager) Commit() error {
	err := txm.DB.Commit().Error
	return err
}

func (txm *TransactionManager) Rollback() {
	txm.DB.Rollback()
}

func (txm *TransactionManager) GetTx() *gorm.DB {
	return txm.DB
}
