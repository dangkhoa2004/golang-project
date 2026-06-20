package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"golang-project/internal/core"
	"golang-project/internal/repository"
)

type UpdateProfileInput struct {
	FullName    string
	Phone       string
	Headline    string
	Bio         string
	Skills      string
	GithubURL   string
	LinkedinURL string
	WebsiteURL  string
}

type UserService interface {
	Register(user *core.User) error
	Login(email, password string) (string, error)
	GetProfileByID(userID uint) (*core.User, error)
	UpdateProfile(userID uint, input *UpdateProfileInput) error
	ChangePassword(userID uint, oldPassword, newPassword string) error
}

type userService struct {
	userRepo repository.UserRepository
	jwtKey   []byte
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
		jwtKey:   []byte("chuoi-bi-mat-sieu-kho-doan-cua-ban"), // Thực tế nên dùng file .env
	}
}

func (s *userService) Register(user *core.User) error {
	existingUser, _ := s.userRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email đã được sử dụng")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("lỗi khi mã hóa mật khẩu")
	}
	user.PasswordHash = string(hashedPassword)

	// Khởi tạo Profile và Streak mặc định
	user.Profile = &core.UserProfile{
		Headline: "Học viên mới",
		Skills:   "[]",
	}
	user.Streak = &core.UserStreak{
		CurrentStreak: 0,
	}

	return s.userRepo.Create(user)
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("sai email hoặc mật khẩu")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("sai email hoặc mật khẩu")
	}

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
	return s.userRepo.FindProfileWithDetails(userID)
}

func (s *userService) UpdateProfile(userID uint, input *UpdateProfileInput) error {
	user := &core.User{
		ID:       userID,
		FullName: input.FullName,
		Phone:    input.Phone,
	}

	profile := &core.UserProfile{
		UserID:      userID,
		Headline:    input.Headline,
		Bio:         input.Bio,
		Skills:      input.Skills,
		GithubURL:   input.GithubURL,
		LinkedinURL: input.LinkedinURL,
		WebsiteURL:  input.WebsiteURL,
	}

	return s.userRepo.UpdateUserAndProfile(user, profile)
}

func (s *userService) ChangePassword(userID uint, oldPassword, newPassword string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("không tìm thấy tài khoản")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
		return errors.New("mật khẩu hiện tại không chính xác")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("lỗi mã hóa mật khẩu")
	}

	return s.userRepo.UpdatePassword(userID, string(hashedPassword))
}
