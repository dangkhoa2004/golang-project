// internal/handler/course_handler.go
package handler

import (
	"net/http"

	"golang-project/internal/service"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService service.CourseService
}

func NewCourseHandler(svc service.CourseService) *CourseHandler {
	return &CourseHandler{courseService: svc}
}

// GET /api/courses?page=1&limit=10
func (h *CourseHandler) GetHomeCourses(c *gin.Context) {
	// SỬ DỤNG CÔNG CỤ TỪ api_request.go
	page := GetQueryInt(c, "page", 1)
	limit := GetQueryInt(c, "limit", 10)

	courses, err := h.courseService.GetCoursesForHomePage(page, limit)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Lỗi lấy dữ liệu khóa học")
		return
	}

	SendSuccess(c, http.StatusOK, "Lấy danh sách khóa học thành công", courses)
}
