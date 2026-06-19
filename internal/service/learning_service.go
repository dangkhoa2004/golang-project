package service

import (
	"errors"
	"golang-project/internal/core"
	"golang-project/internal/repository"
)

type LearningService interface {
	EnrollStudent(studentID uint, courseID uint) error
}

type learningService struct {
	learningRepo repository.LearningRepository
}

func NewLearningService(repo repository.LearningRepository) LearningService {
	return &learningService{learningRepo: repo}
}

func (s *learningService) EnrollStudent(studentID uint, courseID uint) error {
	// Thực tế bạn sẽ cần kiểm tra xem sinh viên đã enroll chưa để tránh lỗi trùng lặp
	enrollment := &core.Enrollment{
		StudentID:       studentID,
		CourseID:        courseID,
		ProgressPercent: 0,
		Status:          "learning",
	}

	err := s.learningRepo.CreateEnrollment(enrollment)
	if err != nil {
		return errors.New("không thể cấp quyền học khóa học này")
	}
	return nil
}
