package core

import "time"

// Bảng orders (Đơn hàng)
type Order struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID     uint      `gorm:"not null" json:"student_id"`
	TotalAmount   float64   `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	Status        string    `gorm:"type:varchar(50);default:'completed'" json:"status"`
	PaymentMethod string    `gorm:"type:varchar(50)" json:"payment_method"`
	CreatedAt     time.Time `gorm:"default:current_timestamp" json:"created_at"`

	Items []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

// Bảng order_items (Chi tiết đơn hàng)
type OrderItem struct {
	ID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID  uint    `gorm:"not null" json:"order_id"`
	CourseID uint    `gorm:"not null" json:"course_id"`
	Price    float64 `gorm:"type:decimal(10,2);not null" json:"price"`

	Course *Course `gorm:"foreignKey:CourseID" json:"course,omitempty"`
}

// Bảng carts (Giỏ hàng)
type Cart struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID uint `gorm:"uniqueIndex;not null" json:"student_id"`

	Items []CartItem `gorm:"foreignKey:CartID" json:"items"`
}

// Bảng cart_items (Sản phẩm trong giỏ)
type CartItem struct {
	ID       uint `gorm:"primaryKey;autoIncrement" json:"id"`
	CartID   uint `gorm:"not null" json:"cart_id"`
	CourseID uint `gorm:"not null" json:"course_id"`

	Course *Course `gorm:"foreignKey:CourseID" json:"course,omitempty"`
}

// Bảng coupons (Mã giảm giá)
type Coupon struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Code            string    `gorm:"type:varchar(50);unique;not null" json:"code"`
	DiscountPercent int       `gorm:"not null" json:"discount_percent"`
	MaxUsage        int       `gorm:"default:100" json:"max_usage"`
	UsedCount       int       `gorm:"default:0" json:"used_count"`
	ExpiryDate      time.Time `json:"expiry_date"`
	IsActive        bool      `gorm:"type:tinyint(1);default:1" json:"is_active"`
}
