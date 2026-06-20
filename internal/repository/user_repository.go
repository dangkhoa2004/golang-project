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
	UpdateUserAndProfile(user *core.User, profile *core.UserProfile) error
	UpdatePassword(userID uint, newPasswordHash string) error
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
	return r.db.Save(profile).Error
}

func (r *userRepository) FindProfileWithDetails(id uint) (*core.User, error) {
	var user core.User
	err := r.db.Preload("Profile").Preload("Streak").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Transaction: Cập nhật 2 bảng cùng lúc
func (r *userRepository) UpdateUserAndProfile(user *core.User, profile *core.UserProfile) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(user).Updates(core.User{FullName: user.FullName, Phone: user.Phone}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", profile.UserID).Updates(profile).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *userRepository) UpdatePassword(userID uint, newPasswordHash string) error {
	return r.db.Model(&core.User{}).Where("id = ?", userID).Update("password_hash", newPasswordHash).Error
}
