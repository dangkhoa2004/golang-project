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

// Struct định nghĩa dữ liệu đầu vào (Payload)
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

// POST /api/register
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest

	// Sử dụng "Công cụ chiết xuất" từ api_request.go
	// Hàm BindJSON sẽ tự lo việc kiểm tra lỗi định dạng và trả về error nếu có
	if err := BindJSON(c, &req); err != nil {
		// Sử dụng "Khuôn đúc kết quả" từ api_response.go
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user := &core.User{
		Email:        req.Email,
		PasswordHash: req.Password,
		FullName:     req.FullName,
		RoleID:       req.RoleID,
	}

	// Gọi logic nghiệp vụ ở tầng Service
	if err := h.userService.Register(user); err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	SendSuccess(c, http.StatusOK, "Đăng ký tài khoản thành công", gin.H{
		"user_id": user.ID,
		"email":   user.Email,
	})
}

// POST /api/login
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest

	// Gom gọn việc bóc tách JSON vào 1 dòng code duy nhất
	if err := BindJSON(c, &req); err != nil {
		SendError(c, http.StatusBadRequest, "Vui lòng cung cấp đầy đủ email đúng định dạng và mật khẩu")
		return
	}

	// Xử lý đăng nhập và sinh JWT
	token, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		SendError(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Trả về Token theo khuôn chuẩn
	SendSuccess(c, http.StatusOK, "Đăng nhập thành công", gin.H{"token": token})
}

// GET /api/user/profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	// Dùng con dao rọc giấy tuyệt vời bạn đã viết ở api_request.go
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		SendError(c, http.StatusUnauthorized, "Không xác định được danh tính")
		return
	}

	// Gọi Service để tìm user theo userID
	user, err := h.userService.GetProfileByID(userID)
	if err != nil {
		SendError(c, http.StatusNotFound, "Không tìm thấy thông tin người dùng")
		return
	}

	// Không bao giờ trả về PasswordHash cho Frontend nhé!
	user.PasswordHash = ""

	SendSuccess(c, http.StatusOK, "Lấy thông tin cá nhân thành công", user)
}
