package mysql

import (
	"gorm.io/gorm"

	"github.com/ktakenaka/go-random/app/domain/entity"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) UpdateOrCreate(body map[string]interface{}) (*entity.User, error) {
	var user entity.User
	err := r.DB.Where(
		entity.User{
			GoogleSub: body["sub"].(string),
		},
	).Assign(
		entity.User{
			Email: body["email"].(string),
		},
	).FirstOrCreate(&user).Error

	return &user, err
}
