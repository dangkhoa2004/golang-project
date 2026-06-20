package service

import (
	"golang-project/internal/core"
	"golang-project/internal/repository"
)

type CourseService interface {
	GetCoursesForHomePage(page, limit int) ([]core.Course, error)
	GetCourseDetail(courseID uint) (*core.Course, error)
	GetCourseCurriculum(courseID uint) (interface{}, error)
	GetCategories() ([]core.Category, error)
	SearchCourses(search string, categoryID uint, sortBy string, page, limit int) ([]core.Course, error)
}

type courseService struct {
	courseRepo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{courseRepo: repo}
}

func (s *courseService) GetCoursesForHomePage(page, limit int) ([]core.Course, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit
	return s.courseRepo.FindAll(limit, offset)
}

func (s *courseService) GetCourseDetail(courseID uint) (*core.Course, error) {
	return s.courseRepo.FindByIDWithDetails(courseID)
}

func (s *courseService) GetCourseCurriculum(courseID uint) (interface{}, error) {
	return nil, nil // TODO: Kết nối với hàm lấy Lộ trình bài học
}

func (s *courseService) GetCategories() ([]core.Category, error) {
	return s.courseRepo.FindCategories()
}

func (s *courseService) SearchCourses(search string, categoryID uint, sortBy string, page, limit int) ([]core.Course, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit
	return s.courseRepo.FindCoursesWithFilters(search, categoryID, sortBy, limit, offset)
}
