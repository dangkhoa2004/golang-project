package config

import (
	"log"
	"time" // Thêm package time để dùng hàm Sleep

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/eduplatform?charset=utf8mb4&parseTime=True&loc=Local"

	var db *gorm.DB
	var err error

	// Sử dụng vòng lặp vô hạn cho đến khi kết nối thành công
	for {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("❌ Không thể kết nối CSDL: %v. Đang thử lại sau 5 giây...", err)
			time.Sleep(5 * time.Second) // Nghỉ 5 giây trước khi lặp lại
			continue
		}

		// Thoát khỏi vòng lặp nếu err == nil (kết nối thành công)
		break
	}

	log.Println("✅ Kết nối CSDL thành công!")
	return db
}
