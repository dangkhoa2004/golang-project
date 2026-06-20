package handler

import (
	"net/http"

	"golang-project/internal/core"
	"golang-project/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{userService: svc}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	RoleID   uint   `json:"role_id" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfileRequest struct {
	FullName    string `json:"full_name" binding:"required"`
	Phone       string `json:"phone"`
	Headline    string `json:"headline"`
	Bio         string `json:"bio"`
	Skills      string `json:"skills"`
	GithubURL   string `json:"github_url"`
	LinkedinURL string `json:"linkedin_url"`
	WebsiteURL  string `json:"website_url"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// POST /api/auth/register
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := BindJSON(c, &req); err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user := &core.User{
		Email: req.Email, PasswordHash: req.Password,
		FullName: req.FullName, RoleID: req.RoleID,
	}

	if err := h.userService.Register(user); err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	SendSuccess(c, http.StatusOK, "Đăng ký tài khoản thành công", gin.H{"user_id": user.ID, "email": user.Email})
}

// POST /api/auth/login
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := BindJSON(c, &req); err != nil {
		SendError(c, http.StatusBadRequest, "Vui lòng cung cấp đầy đủ email và mật khẩu")
		return
	}

	token, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		SendError(c, http.StatusUnauthorized, err.Error())
		return
	}
	SendSuccess(c, http.StatusOK, "Đăng nhập thành công", gin.H{"token": token})
}

// GET /api/users/profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		SendError(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.userService.GetProfileByID(userID)
	if err != nil {
		SendError(c, http.StatusNotFound, "Không tìm thấy thông tin người dùng")
		return
	}
	user.PasswordHash = "" // Xóa password trước khi gửi về
	SendSuccess(c, http.StatusOK, "Lấy thông tin cá nhân thành công", user)
}

// PUT /api/users/profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, _ := GetUserIDFromContext(c)
	var req UpdateProfileRequest
	if err := BindJSON(c, &req); err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	input := &service.UpdateProfileInput{
		FullName: req.FullName, Phone: req.Phone, Headline: req.Headline,
		Bio: req.Bio, Skills: req.Skills, GithubURL: req.GithubURL,
		LinkedinURL: req.LinkedinURL, WebsiteURL: req.WebsiteURL,
	}

	if err := h.userService.UpdateProfile(userID, input); err != nil {
		SendError(c, http.StatusInternalServerError, "Lỗi cập nhật hồ sơ")
		return
	}
	SendSuccess(c, http.StatusOK, "Cập nhật hồ sơ thành công", nil)
}

// PUT /api/users/change-password
func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID, _ := GetUserIDFromContext(c)
	var req ChangePasswordRequest
	if err := BindJSON(c, &req); err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userService.ChangePassword(userID, req.OldPassword, req.NewPassword); err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	SendSuccess(c, http.StatusOK, "Đổi mật khẩu thành công", nil)
}

// CÁC HÀM PLACEHOLDER (Để Router không báo lỗi)
func (h *UserHandler) Logout(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Đăng xuất thành công", nil)
}
func (h *UserHandler) RefreshToken(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Token refreshed", nil)
}
func (h *UserHandler) GetNotifications(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Danh sách thông báo", []string{})
}
