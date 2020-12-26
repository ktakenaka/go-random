package repository

import (
	"gorm.io/gorm"

	"github.com/ktakenaka/go-random/backend/pkg/infra/database"
)

type DBConnection interface {
	Session() *gorm.DB
	Begin() *database.TX // TODO: TXConnection使えるようにリファクタ
}

type TXConnection interface {
	Session() *gorm.DB
	Commit() error
	Rollback() error
}
