package service

import (
	"errors"
	"time"

	"golang-project/internal/core"
	"golang-project/internal/repository"
)

type OrderService interface {
	Checkout(studentID uint, courseIDs []uint, paymentMethod string) (*core.Order, error)
	GetUserOrders(studentID uint) ([]core.Order, error)

	// Các hàm mới
	GetCart(studentID uint) (*core.Cart, error)
	AddToCart(studentID, courseID uint) error
	RemoveFromCart(studentID, courseID uint) error
	ApplyCoupon(code string) (*core.Coupon, error)
}

type orderService struct {
	orderRepo  repository.OrderRepository
	courseRepo repository.CourseRepository
}

func NewOrderService(orderRepo repository.OrderRepository, courseRepo repository.CourseRepository) OrderService {
	return &orderService{orderRepo: orderRepo, courseRepo: courseRepo}
}

func (s *orderService) Checkout(studentID uint, courseIDs []uint, paymentMethod string) (*core.Order, error) {
	if len(courseIDs) == 0 {
		return nil, errors.New("giỏ hàng trống")
	}

	var totalAmount float64
	var orderItems []core.OrderItem

	for _, courseID := range courseIDs {
		course, err := s.courseRepo.FindByIDWithDetails(courseID)
		if err != nil {
			return nil, errors.New("khóa học không tồn tại")
		}

		totalAmount += course.CurrentPrice
		orderItems = append(orderItems, core.OrderItem{
			CourseID: course.ID,
			Price:    course.CurrentPrice,
		})
	}

	order := &core.Order{
		StudentID:     studentID,
		TotalAmount:   totalAmount,
		Status:        "completed",
		PaymentMethod: paymentMethod,
	}

	if err := s.orderRepo.CreateOrderWithItems(order, orderItems); err != nil {
		return nil, errors.New("lỗi khi tạo đơn hàng")
	}

	return order, nil
}

func (s *orderService) GetUserOrders(studentID uint) ([]core.Order, error) {
	return s.orderRepo.FindOrdersByUserID(studentID)
}

func (s *orderService) GetCart(studentID uint) (*core.Cart, error) {
	return s.orderRepo.GetCartByStudentID(studentID)
}

func (s *orderService) AddToCart(studentID, courseID uint) error {
	cart, err := s.orderRepo.GetCartByStudentID(studentID)
	if err != nil {
		return errors.New("lỗi truy xuất giỏ hàng")
	}

	if s.orderRepo.CheckCourseInCart(cart.ID, courseID) {
		return errors.New("khóa học này đã có trong giỏ hàng của bạn")
	}

	item := &core.CartItem{
		CartID:   cart.ID,
		CourseID: courseID,
	}
	return s.orderRepo.AddCartItem(item)
}

func (s *orderService) RemoveFromCart(studentID, courseID uint) error {
	cart, err := s.orderRepo.GetCartByStudentID(studentID)
	if err != nil {
		return errors.New("lỗi truy xuất giỏ hàng")
	}
	return s.orderRepo.RemoveCartItem(cart.ID, courseID)
}

func (s *orderService) ApplyCoupon(code string) (*core.Coupon, error) {
	coupon, err := s.orderRepo.FindCouponByCode(code)
	if err != nil {
		return nil, errors.New("mã giảm giá không hợp lệ hoặc không tồn tại")
	}

	if coupon.UsedCount >= coupon.MaxUsage {
		return nil, errors.New("mã giảm giá đã hết lượt sử dụng")
	}

	if time.Now().After(coupon.ExpiryDate) {
		return nil, errors.New("mã giảm giá đã hết hạn")
	}

	return coupon, nil
}
