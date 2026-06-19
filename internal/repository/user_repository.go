package repository

import (
	"golang-project/internal/core"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *core.User) error
	FindByID(id uint) (*core.User, error)
	FindByEmail(email string) (*core.User, error)
	UpdateProfile(profile *core.UserProfile) error
	FindProfileWithDetails(id uint) (*core.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *core.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id uint) (*core.User, error) {
	var user core.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*core.User, error) {
	var user core.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateProfile(profile *core.UserProfile) error {
	// Ghi đè hoặc tạo mới profile (Upsert)
	return r.db.Save(profile).Error
}

// Lấy thông tin User kèm theo Profile và Streak
func (r *userRepository) FindProfileWithDetails(id uint) (*core.User, error) {
	var user core.User
	err := r.db.Preload("Profile").Preload("Streak").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
