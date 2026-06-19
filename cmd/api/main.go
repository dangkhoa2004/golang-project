package main

import (
	"log"

	"golang-project/internal/config"
	"golang-project/internal/router"
)

func main() {
	// 1. Kết nối Database
	db := config.ConnectDB()

	// 2. Khởi tạo Router và các thành phần
	r := router.Setup(db)

	// 3. Chạy Server
	log.Println("🚀 Server đang chạy tại http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Lỗi khởi động server: %v", err)
	}
}
