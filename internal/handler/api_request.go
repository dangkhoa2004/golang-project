package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 1. Gói gọn logic Bind JSON và xuất ra câu lỗi chung dễ hiểu
func BindJSON(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(req); err != nil {
		// Trong thực tế, bạn có thể custom thư viện validator ở đây để báo lỗi chi tiết hơn
		return errors.New("dữ liệu đầu vào không hợp lệ hoặc thiếu trường bắt buộc")
	}
	return nil
}

// 2. Lấy Query Parameter dạng số (thường dùng cho phân trang: ?page=1&limit=10)
func GetQueryInt(c *gin.Context, key string, defaultValue int) int {
	valStr := c.DefaultQuery(key, strconv.Itoa(defaultValue))
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultValue // Nếu frontend truyền linh tinh (VD: ?page=abc), trả về mặc định
	}
	return val
}

// 3. Lấy URI Parameter dạng số (VD: /api/courses/:id)
func GetParamInt(c *gin.Context, key string) (uint, error) {
	valStr := c.Param(key)
	val, err := strconv.Atoi(valStr)
	if err != nil || val <= 0 {
		return 0, errors.New("tham số URL không hợp lệ (phải là số dương)")
	}
	return uint(val), nil
}

// 4. Lấy ID người dùng từ Context (Do Auth Middleware truyền vào)
func GetUserIDFromContext(c *gin.Context) (uint, error) {
	val, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("không tìm thấy thông tin người dùng, vui lòng đăng nhập")
	}

	userID, ok := val.(uint)
	if !ok {
		return 0, errors.New("dữ liệu người dùng bị lỗi định dạng")
	}

	return userID, nil
}
