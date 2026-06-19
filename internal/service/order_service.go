package service

import (
	"errors"
	"golang-project/internal/core"
	"golang-project/internal/repository"
)

type OrderService interface {
	Checkout(studentID uint, courseIDs []uint, paymentMethod string) (*core.Order, error)
}

type orderService struct {
	orderRepo  repository.OrderRepository
	courseRepo repository.CourseRepository // Tiêm thêm repo này để lấy giá tiền
}

func NewOrderService(orderRepo repository.OrderRepository, courseRepo repository.CourseRepository) OrderService {
	return &orderService{
		orderRepo:  orderRepo,
		courseRepo: courseRepo,
	}
}

func (s *orderService) Checkout(studentID uint, courseIDs []uint, paymentMethod string) (*core.Order, error) {
	if len(courseIDs) == 0 {
		return nil, errors.New("giỏ hàng trống")
	}

	var totalAmount float64
	var orderItems []core.OrderItem

	// Lặp qua từng khóa học để tính tổng tiền và tạo OrderItem
	for _, id := range courseIDs {
		course, err := s.courseRepo.FindByIDWithDetails(id)
		if err != nil {
			return nil, errors.New("khóa học không tồn tại")
		}

		totalAmount += course.CurrentPrice

		orderItems = append(orderItems, core.OrderItem{
			CourseID: course.ID,
			Price:    course.CurrentPrice, // Lưu lại giá tại thời điểm mua (tránh việc sau này khóa học tăng giá làm sai lịch sử)
		})
	}

	// Tạo đối tượng Order
	order := &core.Order{
		StudentID:     studentID,
		TotalAmount:   totalAmount,
		Status:        "completed", // Thực tế sẽ là "pending" chờ thanh toán Momo/VNPay
		PaymentMethod: paymentMethod,
	}

	// Gọi Transaction ở Repo để lưu cả đơn hàng lẫn chi tiết
	err := s.orderRepo.CreateOrderWithItems(order, orderItems)
	if err != nil {
		return nil, errors.New("lỗi khi tạo đơn hàng")
	}

	return order, nil
}
