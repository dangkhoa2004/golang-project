package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Khóa bí mật (phải khớp với khóa bên file user_service.go)
var jwtKey = []byte("chuoi-bi-mat-sieu-kho-doan-cua-ban")

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Lấy Token từ Header: "Authorization: Bearer <token>"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Vui lòng đăng nhập"})
			c.Abort() // Chặn không cho đi tiếp vào Handler
			return
		}

		// 2. Tách chữ "Bearer " ra để lấy đúng chuỗi token
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// 3. Giải mã và xác thực Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("phương thức ký không hợp lệ")
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ hoặc đã hết hạn"})
			c.Abort()
			return
		}

		// 4. Bóc tách dữ liệu (Claims) và lưu vào Context của Gin để Handler dùng
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Ép kiểu user_id về uint
			userID := uint(claims["user_id"].(float64))
			c.Set("user_id", userID)
			c.Set("role_id", uint(claims["role_id"].(float64)))
		}

		c.Next() // Cho phép đi tiếp vào Handler
	}
}
