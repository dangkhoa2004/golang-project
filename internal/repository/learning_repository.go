package repository

import (
	"golang-project/internal/core"

	"gorm.io/gorm"
)

type LearningRepository interface {
	CreateEnrollment(enrollment *core.Enrollment) error
	GetMyCourses(studentID uint) ([]core.Enrollment, error)
	CheckEnrollment(studentID, courseID uint) bool
	GetLessonProgress(enrollmentID, lessonID uint) (*core.LessonProgress, error)
	UpsertProgress(progress *core.LessonProgress) error
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

func (r *learningRepository) GetMyCourses(studentID uint) ([]core.Enrollment, error) {
	var enrollments []core.Enrollment

	// SỬA Ở ĐÂY: Thêm Preload("Course.Teacher") để lấy sâu vào thông tin giảng viên của khóa học
	err := r.db.
		Preload("Course").
		Preload("Course.Teacher").
		Where("student_id = ?", studentID).
		Find(&enrollments).Error

	return enrollments, err
}

func (r *learningRepository) CheckEnrollment(studentID, courseID uint) bool {
	var count int64
	r.db.Model(&core.Enrollment{}).Where("student_id = ? AND course_id = ?", studentID, courseID).Count(&count)
	return count > 0
}

func (r *learningRepository) GetLessonProgress(enrollmentID, lessonID uint) (*core.LessonProgress, error) {
	var progress core.LessonProgress
	err := r.db.Where("enrollment_id = ? AND lesson_id = ?", enrollmentID, lessonID).First(&progress).Error
	if err != nil {
		return nil, err
	}
	return &progress, nil
}

func (r *learningRepository) UpsertProgress(progress *core.LessonProgress) error {
	// GORM Save() tự động INSERT nếu chưa có ID, hoặc UPDATE nếu đã tồn tại
	return r.db.Save(progress).Error
}

func (r *learningRepository) GetDiscussionsByLesson(lessonID uint) ([]core.Discussion, error) {
	var discussions []core.Discussion
	err := r.db.Where("lesson_id = ?", lessonID).Order("created_at desc").Find(&discussions).Error
	return discussions, err
}
