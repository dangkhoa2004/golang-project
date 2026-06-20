package service

import (
	"errors"
	"golang-project/internal/core"
	"golang-project/internal/repository"
)

type LearningService interface {
	GetMyCourses(studentID uint) ([]core.Enrollment, error)
	GetLessonSecureContent(studentID, courseID, lessonID uint) (interface{}, error) // Đổi kiểu trả về tạm thời
	TrackVideoProgress(studentID, courseID, lessonID uint, watchedSeconds int) error
	MarkLessonComplete(studentID, courseID, lessonID uint) error
}

type learningService struct {
	learningRepo repository.LearningRepository
	courseRepo   repository.CourseRepository
}

func NewLearningService(lr repository.LearningRepository, cr repository.CourseRepository) LearningService {
	return &learningService{learningRepo: lr, courseRepo: cr}
}

func (s *learningService) GetMyCourses(studentID uint) ([]core.Enrollment, error) {
	return s.learningRepo.GetMyCourses(studentID)
}

func (s *learningService) GetLessonSecureContent(studentID, courseID, lessonID uint) (interface{}, error) {
	hasAccess := s.learningRepo.CheckEnrollment(studentID, courseID)
	if !hasAccess {
		return nil, errors.New("bạn chưa mua khóa học này, vui lòng thanh toán để xem")
	}
	// TODO: Tạo hàm GetLessonByID trong courseRepo để trả về Video URL
	return nil, nil
}

func (s *learningService) TrackVideoProgress(studentID, courseID, lessonID uint, watchedSeconds int) error {
	// Giả lập lấy EnrollmentID (bạn cần thêm GetEnrollmentByStudentAndCourse vào LearningRepo sau này)
	// Tạm thời bỏ qua bước check EnrollmentID chi tiết để code build thành công
	progress := &core.LessonProgress{
		EnrollmentID:      1, // Cần thay bằng ID thật
		LessonID:          lessonID,
		LastWatchedSecond: watchedSeconds,
	}
	return s.learningRepo.UpsertProgress(progress)
}

func (s *learningService) MarkLessonComplete(studentID, courseID, lessonID uint) error {
	progress := &core.LessonProgress{
		EnrollmentID: 1, // Cần thay bằng ID thật
		LessonID:     lessonID,
		IsCompleted:  true,
	}
	return s.learningRepo.UpsertProgress(progress)
}
