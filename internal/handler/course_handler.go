package handler

import (
	"golang-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService service.CourseService
}

func NewCourseHandler(svc service.CourseService) *CourseHandler {
	return &CourseHandler{courseService: svc}
}

// GET /api/categories
func (h *CourseHandler) GetCategories(c *gin.Context) {
	categories, err := h.courseService.GetCategories()
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Lỗi lấy danh mục")
		return
	}
	SendSuccess(c, http.StatusOK, "Thành công", categories)
}

// GET /api/courses?search=&category_id=&sort_by=&page=&limit=
func (h *CourseHandler) GetCourses(c *gin.Context) {
	search := c.Query("search")
	sortBy := c.Query("sort_by")
	categoryID := GetQueryInt(c, "category_id", 0)
	page := GetQueryInt(c, "page", 1)
	limit := GetQueryInt(c, "limit", 10)

	courses, err := h.courseService.SearchCourses(search, uint(categoryID), sortBy, page, limit)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Lỗi lấy dữ liệu khóa học")
		return
	}
	SendSuccess(c, http.StatusOK, "Thành công", courses)
}

// GET /api/courses/:courseId
func (h *CourseHandler) GetCourseDetail(c *gin.Context) {
	id, err := GetParamInt(c, "courseId")
	if err != nil {
		SendError(c, http.StatusBadRequest, "ID khóa học không hợp lệ")
		return
	}

	course, err := h.courseService.GetCourseDetail(id)
	if err != nil {
		SendError(c, http.StatusNotFound, "Không tìm thấy khóa học")
		return
	}
	SendSuccess(c, http.StatusOK, "Thành công", course)
}

// GET /api/courses/:courseId/reviews
func (h *CourseHandler) GetCourseReviews(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Danh sách đánh giá", []string{})
}
