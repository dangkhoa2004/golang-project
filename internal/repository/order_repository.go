package repository

import (
	"golang-project/internal/core"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrderWithItems(order *core.Order, items []core.OrderItem) error
	FindOrderByID(id uint) (*core.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

// Ứng dụng Transaction để đảm bảo tính toàn vẹn dữ liệu
func (r *orderRepository) CreateOrderWithItems(order *core.Order, items []core.OrderItem) error {
	// Bắt đầu một Transaction
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Tạo Order trước để lấy được ID
		if err := tx.Create(order).Error; err != nil {
			return err // Return err sẽ tự động Rollback
		}

		// 2. Gán OrderID vừa được tạo cho từng Item và lưu vào bảng order_items
		for i := range items {
			items[i].OrderID = order.ID
		}

		if err := tx.Create(&items).Error; err != nil {
			return err // Return err sẽ tự động Rollback
		}

		// Trả về nil nghĩa là mọi thứ thành công, Transaction sẽ tự động Commit
		return nil
	})
}

func (r *orderRepository) FindOrderByID(id uint) (*core.Order, error) {
	var order core.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
