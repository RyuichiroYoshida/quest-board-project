package infrastructure

import (
	"time"

	"github.com/RyuichiroYoshida/quest-board-project/internal/auth/domain"
	"github.com/RyuichiroYoshida/quest-board-project/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *domain.User) error
	ReadUser(id string) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) CreateUser(user *domain.User) error {
	// domain.User -> models.User への変換
	mUser := models.User{
		Id:        user.Id,
		Name:      user.Name,
		Avatar:    user.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return r.db.Create(&mUser).Error
}

func (r *authRepository) ReadUser(id string) (*domain.User, error) {
	var mUser models.User
	if err := r.db.Where("id = ?", id).First(&mUser).Error; err != nil {
		return nil, err
	}

	// models.User -> domain.User への変換
	user := &domain.User{
		Id:     mUser.Id,
		Name:   mUser.Name,
		Avatar: mUser.Avatar,
	}

	return user, nil
}

func (r *authRepository) UpdateUser(user *domain.User) error {
	// domain.User -> models.User への変換
	mUser := models.User{
		Id:        user.Id,
		Name:      user.Name,
		Avatar:    user.Avatar,
		UpdatedAt: time.Now(),
	}

	return r.db.Save(&mUser).Error
}

func (r *authRepository) DeleteUser(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.User{}).Error
}
