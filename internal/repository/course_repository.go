package repository

import (
	"golang-project/internal/core"

	"gorm.io/gorm"
)

type CourseRepository interface {
	FindAll(limit int, offset int) ([]core.Course, error)
	FindByIDWithDetails(id uint) (*core.Course, error)
	FindLessonsByChapterID(chapterID uint) ([]core.Lesson, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) FindAll(limit int, offset int) ([]core.Course, error) {
	var courses []core.Course
	// Lấy danh sách khóa học, tự động join với bảng Tags
	err := r.db.Preload("Tags").Limit(limit).Offset(offset).Find(&courses).Error
	return courses, err
}

func (r *courseRepository) FindByIDWithDetails(id uint) (*core.Course, error) {
	var course core.Course
	// Preload sâu (Eager Loading): Lấy Course -> Kèm Categories -> Kèm Tags
	// Lưu ý: Để Preload hoạt động, cấu trúc Struct ở tầng Core cần khai báo Relationship (Tags []Tag, Chapter []Chapter...)
	err := r.db.Preload("Tags").Where("id = ?", id).First(&course).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) FindLessonsByChapterID(chapterID uint) ([]core.Lesson, error) {
	var lessons []core.Lesson
	err := r.db.Where("chapter_id = ?", chapterID).Order("order_index asc").Find(&lessons).Error
	return lessons, err
}
