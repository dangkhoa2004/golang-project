package service

import (
	"golang-project/internal/core"
	"golang-project/internal/repository"
)

type CourseService interface {
	GetCoursesForHomePage(page, limit int) ([]core.Course, error)
	GetCourseDetail(courseID uint) (*core.Course, error)
}

type courseService struct {
	courseRepo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{courseRepo: repo}
}

func (s *courseService) GetCoursesForHomePage(page, limit int) ([]core.Course, error) {
	// Tính toán offset cho phân trang (Pagination)
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	return s.courseRepo.FindAll(limit, offset)
}

func (s *courseService) GetCourseDetail(courseID uint) (*core.Course, error) {
	// Ở đây bạn có thể thêm logic kiểm tra xem khóa học có đang bị ẩn (status != 'published') không
	return s.courseRepo.FindByIDWithDetails(courseID)
}
