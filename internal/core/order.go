package core

import "time"

// Bảng orders
type Order struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID     uint      `gorm:"not null" json:"student_id"`
	TotalAmount   float64   `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	Status        string    `gorm:"type:varchar(50);default:'completed'" json:"status"`
	PaymentMethod string    `gorm:"type:varchar(50)" json:"payment_method"`
	CreatedAt     time.Time `gorm:"default:current_timestamp" json:"created_at"`
}

// Bảng order_items
type OrderItem struct {
	ID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID  uint    `gorm:"not null" json:"order_id"`
	CourseID uint    `gorm:"not null" json:"course_id"`
	Price    float64 `gorm:"type:decimal(10,2);not null" json:"price"`
}
