// internal/handler/api_response.go
package handler

import (
	"github.com/gin-gonic/gin"
)

// Cấu trúc bắt buộc cho mọi API Response
type APIResponse struct {
	RequestFrom string      `json:"request_from"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	Status      string      `json:"status"`
}

// SendSuccess là hàm tiện ích dùng để trả về dữ liệu khi thành công
func SendSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	res := APIResponse{
		RequestFrom: c.Request.URL.Path, // Tự động lấy URL của request hiện tại
		Data:        data,
		Message:     message,
		Status:      "success",
	}
	c.JSON(statusCode, res)
}

// SendError là hàm tiện ích dùng để trả về lỗi
func SendError(c *gin.Context, statusCode int, message string) {
	res := APIResponse{
		RequestFrom: c.Request.URL.Path,
		Data:        nil, // Khi lỗi thì thường không có data
		Message:     message,
		Status:      "error",
	}
	c.JSON(statusCode, res)
}
