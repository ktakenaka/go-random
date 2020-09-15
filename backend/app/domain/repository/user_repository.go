package repository

import (
	"github.com/ktakenaka/go-random/backend/app/domain/entity"
)

type UserRepository interface {
	UpdateOrCreate(body map[string]interface{}) (*entity.User, error)
}
