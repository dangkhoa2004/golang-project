package repository

import (
	"golang-project/internal/core"

	"gorm.io/gorm"
)

type CourseRepository interface {
	FindAll(limit int, offset int) ([]core.Course, error)
	FindByIDWithDetails(id uint) (*core.Course, error)
	FindLessonsByChapterID(chapterID uint) ([]core.Lesson, error)
	FindCategories() ([]core.Category, error)
	FindCoursesWithFilters(search string, categoryID uint, sortBy string, limit, offset int) ([]core.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) FindAll(limit int, offset int) ([]core.Course, error) {
	var courses []core.Course
	// ĐÃ THÊM: Preload("Teacher") để lấy tên giảng viên
	err := r.db.
		Preload("Tags").
		Preload("Teacher").
		Limit(limit).Offset(offset).Find(&courses).Error
	return courses, err
}

func (r *courseRepository) FindByIDWithDetails(id uint) (*core.Course, error) {
	var course core.Course

	// Preload sâu tất cả dữ liệu liên quan
	err := r.db.
		Preload("Tags").
		Preload("Category").
		Preload("Teacher").         // Lấy thông tin User của giảng viên
		Preload("Teacher.Profile"). // Lấy Bio/Headline của giảng viên
		Preload("Chapters", func(db *gorm.DB) *gorm.DB {
			return db.Order("order_index asc") // Sắp xếp Chương theo thứ tự
		}).
		Preload("Chapters.Lessons", func(db *gorm.DB) *gorm.DB {
			return db.Order("order_index asc") // Sắp xếp Bài học theo thứ tự
		}).
		Where("id = ?", id).First(&course).Error

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

func (r *courseRepository) FindCategories() ([]core.Category, error) {
	var categories []core.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *courseRepository) FindCoursesWithFilters(search string, categoryID uint, sortBy string, limit, offset int) ([]core.Course, error) {
	var courses []core.Course

	// SỬA Ở ĐÂY: Thêm Preload("Teacher") vào chuỗi truy vấn
	query := r.db.
		Preload("Tags").
		Preload("Teacher").
		Where("status = ?", "published")

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}
	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	switch sortBy {
	case "priceAsc":
		query = query.Order("current_price asc")
	case "priceDesc":
		query = query.Order("current_price desc")
	case "rating":
		query = query.Order("avg_rating desc")
	default:
		query = query.Order("created_at desc")
	}

	err := query.Limit(limit).Offset(offset).Find(&courses).Error
	return courses, err
}
