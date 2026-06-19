package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"golang-project/internal/core"
	"golang-project/internal/repository"
)

type UserService interface {
	Register(user *core.User) error
	Login(email, password string) (string, error)
	GetProfileByID(userID uint) (*core.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	jwtKey   []byte // Secret key để ký JWT (Trong thực tế nên để ở file .env)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
		jwtKey:   []byte("chuoi-bi-mat-sieu-kho-doan-cua-ban"),
	}
}

func (s *userService) Register(user *core.User) error {
	// 1. Kiểm tra email đã tồn tại chưa
	existingUser, _ := s.userRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email đã được sử dụng")
	}

	// 2. Mã hóa mật khẩu bằng Bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("lỗi khi mã hóa mật khẩu")
	}
	user.PasswordHash = string(hashedPassword)

	// 3. Gọi Repo để lưu vào DB
	return s.userRepo.Create(user)
}

func (s *userService) Login(email, password string) (string, error) {
	// 1. Tìm user theo email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("sai email hoặc mật khẩu")
	}

	// 2. So sánh mật khẩu người dùng nhập vào với mã băm trong DB
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("sai email hoặc mật khẩu")
	}

	// 3. Tạo JWT Token (thời hạn 24 giờ)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role_id": user.RoleID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", errors.New("không thể tạo token đăng nhập")
	}

	return tokenString, nil
}
func (s *userService) GetProfileByID(userID uint) (*core.User, error) {
	// Đổi thành FindProfileWithDetails để GORM tự động JOIN lấy Profile và Streak
	return s.userRepo.FindProfileWithDetails(userID)
}
