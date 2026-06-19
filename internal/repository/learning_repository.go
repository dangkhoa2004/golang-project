package repository

import (
	"golang-project/internal/core"

	"gorm.io/gorm"
)

type LearningRepository interface {
	CreateEnrollment(enrollment *core.Enrollment) error
	UpdateLessonProgress(progress *core.LessonProgress) error
	GetDiscussionsByLesson(lessonID uint) ([]core.Discussion, error)
}

type learningRepository struct {
	db *gorm.DB
}

func NewLearningRepository(db *gorm.DB) LearningRepository {
	return &learningRepository{db: db}
}

func (r *learningRepository) CreateEnrollment(enrollment *core.Enrollment) error {
	return r.db.Create(enrollment).Error
}

func (r *learningRepository) UpdateLessonProgress(progress *core.LessonProgress) error {
	// Dùng Where kết hợp Updates để cập nhật tiến độ theo LessonID và EnrollmentID
	return r.db.Where("enrollment_id = ? AND lesson_id = ?", progress.EnrollmentID, progress.LessonID).
		Updates(progress).Error
}

func (r *learningRepository) GetDiscussionsByLesson(lessonID uint) ([]core.Discussion, error) {
	var discussions []core.Discussion
	err := r.db.Where("lesson_id = ?", lessonID).Order("created_at desc").Find(&discussions).Error
	return discussions, err
}
