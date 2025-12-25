package repo

import (
	"errors"

	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
	"gorm.io/gorm"
)

type PostgresUserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) core.UserRepository {
	return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) CreateUser(user *core.User) error {
	entity := UserEntity{
		Email:    user.Email,
		Password: user.Password,
	}

	result := r.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}

	user.ID = entity.ID
	return nil
}

func (r *PostgresUserRepo) GetUserByEmail(email string) (*core.User, error) {
	var entity UserEntity

	result := r.db.Where("email = ?", email).First(&entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return &core.User{
		ID:       entity.ID,
		Email:    entity.Email,
		Password: entity.Password,
	}, nil
}
