package repository

import (
	"golang-project/internal/core"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrderWithItems(order *core.Order, items []core.OrderItem) error
	FindOrderByID(id uint) (*core.Order, error)
	FindOrdersByUserID(userID uint) ([]core.Order, error)
	GetCartByStudentID(studentID uint) (*core.Cart, error)
	AddCartItem(item *core.CartItem) error
	RemoveCartItem(cartID, courseID uint) error
	CheckCourseInCart(cartID, courseID uint) bool
	FindCouponByCode(code string) (*core.Coupon, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrderWithItems(order *core.Order, items []core.OrderItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		for i := range items {
			items[i].OrderID = order.ID
		}
		if err := tx.Create(&items).Error; err != nil {
			return err
		}
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

func (r *orderRepository) FindOrdersByUserID(userID uint) ([]core.Order, error) {
	var orders []core.Order
	err := r.db.Preload("Items").Preload("Items.Course").Where("student_id = ?", userID).Order("created_at desc").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetCartByStudentID(studentID uint) (*core.Cart, error) {
	var cart core.Cart
	// FirstOrCreate: Tìm giỏ hàng, nếu không có thì tự tạo mới một giỏ trống cho User
	if err := r.db.Preload("Items").Preload("Items.Course").FirstOrCreate(&cart, core.Cart{StudentID: studentID}).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *orderRepository) CheckCourseInCart(cartID, courseID uint) bool {
	var count int64
	r.db.Model(&core.CartItem{}).Where("cart_id = ? AND course_id = ?", cartID, courseID).Count(&count)
	return count > 0
}

func (r *orderRepository) AddCartItem(item *core.CartItem) error {
	return r.db.Create(item).Error
}

func (r *orderRepository) RemoveCartItem(cartID, courseID uint) error {
	return r.db.Where("cart_id = ? AND course_id = ?", cartID, courseID).Delete(&core.CartItem{}).Error
}

func (r *orderRepository) FindCouponByCode(code string) (*core.Coupon, error) {
	var coupon core.Coupon
	if err := r.db.Where("code = ? AND is_active = ?", code, true).First(&coupon).Error; err != nil {
		return nil, err
	}
	return &coupon, nil
}
