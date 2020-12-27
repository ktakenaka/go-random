package mysql

import (
	"gorm.io/gorm"

	"github.com/ktakenaka/go-random/backend/app/domain/entity"
	"github.com/ktakenaka/go-random/backend/pkg/infra/database"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{DB: db.Session()}
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
