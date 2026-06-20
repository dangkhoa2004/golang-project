package handler

import (
	"golang-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(svc service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: svc}
}

type AddCartRequest struct {
	CourseID uint `json:"course_id" binding:"required"`
}

type CouponRequest struct {
	Code string `json:"code" binding:"required"`
}

type CheckoutRequest struct {
	CourseIDs     []uint `json:"course_ids" binding:"required,min=1"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

// GET /api/cart
func (h *OrderHandler) GetCart(c *gin.Context) {
	studentID, _ := GetUserIDFromContext(c)
	cart, err := h.orderService.GetCart(studentID)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Lỗi lấy thông tin giỏ hàng")
		return
	}
	SendSuccess(c, http.StatusOK, "Thành công", cart)
}

// POST /api/cart/add
func (h *OrderHandler) AddToCart(c *gin.Context) {
	studentID, _ := GetUserIDFromContext(c)
	var req AddCartRequest
	if err := BindJSON(c, &req); err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.orderService.AddToCart(studentID, req.CourseID); err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	SendSuccess(c, http.StatusOK, "Đã thêm khóa học vào giỏ", nil)
}

// DELETE /api/cart/remove/:courseId
func (h *OrderHandler) RemoveFromCart(c *gin.Context) {
	studentID, _ := GetUserIDFromContext(c)
	courseID, err := GetParamInt(c, "courseId")
	if err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.orderService.RemoveFromCart(studentID, courseID); err != nil {
		SendError(c, http.StatusInternalServerError, "Lỗi xóa khỏi giỏ hàng")
		return
	}
	SendSuccess(c, http.StatusOK, "Đã xóa khóa học khỏi giỏ", nil)
}

// POST /api/cart/apply-coupon
func (h *OrderHandler) ApplyCoupon(c *gin.Context) {
	var req CouponRequest
	if err := BindJSON(c, &req); err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	coupon, err := h.orderService.ApplyCoupon(req.Code)
	if err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	SendSuccess(c, http.StatusOK, "Áp dụng mã giảm giá thành công", coupon)
}

// POST /api/orders
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CheckoutRequest
	if err := BindJSON(c, &req); err != nil {
		SendError(c, http.StatusBadRequest, "Dữ liệu thanh toán không hợp lệ")
		return
	}

	studentID, _ := GetUserIDFromContext(c)
	order, err := h.orderService.Checkout(studentID, req.CourseIDs, req.PaymentMethod)
	if err != nil {
		SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccess(c, http.StatusOK, "Tạo đơn hàng thành công", order)
}

// GET /api/orders/history
func (h *OrderHandler) GetOrderHistory(c *gin.Context) {
	studentID, _ := GetUserIDFromContext(c)
	orders, err := h.orderService.GetUserOrders(studentID)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Lỗi khi tải lịch sử giao dịch")
		return
	}
	SendSuccess(c, http.StatusOK, "Tải lịch sử giao dịch thành công", orders)
}

// Placeholder
func (h *OrderHandler) ProcessPayment(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Đang xử lý cổng thanh toán...", nil)
}
