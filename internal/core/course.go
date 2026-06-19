package core

import "time"

// Bảng categories
type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Slug string `gorm:"type:varchar(100);unique;not null" json:"slug"`
}

// Bảng tags
type Tag struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}

// Bảng courses
type Course struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	TeacherID   uint   `gorm:"not null" json:"teacher_id"`
	CategoryID  uint   `gorm:"not null" json:"category_id"`
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`

	// SỬA Ở ĐÂY: Thêm column:what_you_will_learn
	WhatYouLearn string `gorm:"column:what_you_will_learn;type:text" json:"what_will_you_learn"`

	ThumbnailURL  string  `gorm:"type:varchar(255)" json:"thumbnail_url"`
	Level         string  `gorm:"type:varchar(50)" json:"level"`
	CourseType    string  `gorm:"type:varchar(50);default:'recorded'" json:"course_type"`
	IsFree        bool    `gorm:"type:tinyint(1);default:0" json:"is_free"`
	OriginalPrice float64 `gorm:"type:decimal(10,2);default:0.00" json:"original_price"`
	CurrentPrice  float64 `gorm:"type:decimal(10,2);default:0.00" json:"current_price"`
	AvgRating     float64 `gorm:"type:decimal(3,2);default:0.00" json:"avg_rating"`
	ReviewCount   int     `gorm:"default:0" json:"review_count"`
	StudentCount  int     `gorm:"default:0" json:"student_count"`

	// SỬA Ở ĐÂY: Thêm column:total_lessons
	TotalLesson int `gorm:"column:total_lessons;default:0" json:"total_lesson"`

	// SỬA Ở ĐÂY: Thêm column:total_duration_seconds
	TotalDuration int `gorm:"column:total_duration_seconds;default:0" json:"total_duration"`

	Status    string    `gorm:"type:varchar(50);default:'published'" json:"status"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`

	// Thiết lập mối quan hệ n-n với bảng tags thông qua bảng trung gian course_tags
	Tags []Tag `gorm:"many2many:course_tags;" json:"tags,omitempty"`
}

// Bảng trung gian course_tags (GORM tự động map nhưng viết ra để tường minh)
type CourseTag struct {
	CourseID uint `gorm:"primaryKey" json:"course_id"`
	TagID    uint `gorm:"primaryKey" json:"tag_id"`
}

// Bảng chapters
type Chapter struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	CourseID   uint   `gorm:"not null" json:"course_id"`
	Title      string `gorm:"type:varchar(255);not null" json:"title"`
	OrderIndex int    `json:"order_index"`
}

// Bảng lessons
type Lesson struct {
	ID              uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ChapterID       uint   `gorm:"not null" json:"chapter_id"`
	Title           string `gorm:"type:varchar(255);not null" json:"title"`
	VideoURL        string `gorm:"type:varchar(255)" json:"video_url"`
	DurationSeconds int    `json:"duration_seconds"`
	OrderIndex      int    `json:"order_index"`
}

// Bảng lesson_attachments
type LessonAttachment struct {
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	LessonID   uint    `gorm:"not null" json:"lesson_id"`
	FileName   string  `gorm:"type:varchar(255);not null" json:"file_name"`
	FileURL    string  `gorm:"type:varchar(255);not null" json:"file_url"`
	FileSizeMB float64 `gorm:"type:decimal(5,2)" json:"file_size_mb"`
}

// Bảng wishlists
type Wishlist struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	CourseID  uint      `gorm:"not null" json:"course_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
}
