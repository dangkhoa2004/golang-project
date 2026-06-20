package core

import "time"

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"-"` // Ẩn ID
	Name string `gorm:"type:varchar(100)" json:"name"`
	Slug string `gorm:"type:varchar(100)" json:"-"`
}

type Tag struct {
	ID   uint   `gorm:"primaryKey" json:"-"`
	Name string `gorm:"type:varchar(50)" json:"name"`
}

type Course struct {
	ID            uint      `gorm:"primaryKey" json:"-"` // Ẩn ID khóa học
	TeacherID     uint      `gorm:"not null" json:"-"`   // Ẩn ID
	CategoryID    uint      `gorm:"not null" json:"-"`   // Ẩn ID
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	WhatYouLearn  string    `gorm:"column:what_you_will_learn" json:"what_will_you_learn"`
	ThumbnailURL  string    `json:"thumbnail_url"`
	Level         string    `json:"level"`
	CourseType    string    `json:"course_type"`
	IsFree        bool      `json:"is_free"`
	OriginalPrice float64   `json:"original_price"`
	CurrentPrice  float64   `json:"current_price"`
	AvgRating     float64   `json:"avg_rating"`
	ReviewCount   int       `json:"review_count"`
	StudentCount  int       `json:"student_count"`
	TotalLesson   int       `json:"total_lesson"`
	TotalDuration int       `json:"total_duration"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`

	Tags     []Tag     `gorm:"many2many:course_tags;" json:"tags,omitempty"`
	Teacher  *User     `gorm:"foreignKey:TeacherID" json:"teacher,omitempty"`
	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Chapters []Chapter `gorm:"foreignKey:CourseID" json:"chapters,omitempty"`
}

type Chapter struct {
	ID         uint     `gorm:"primaryKey" json:"-"` // Ẩn ID chương
	CourseID   uint     `gorm:"not null" json:"-"`   // Ẩn ID
	Title      string   `json:"title"`
	OrderIndex int      `json:"order_index"`
	Lessons    []Lesson `gorm:"foreignKey:ChapterID" json:"lessons,omitempty"`
}

type Lesson struct {
	ID              uint   `gorm:"primaryKey" json:"-"` // Ẩn ID bài học
	ChapterID       uint   `gorm:"not null" json:"-"`   // Ẩn ID
	Title           string `json:"title"`
	VideoURL        string `json:"video_url,omitempty"` // Chỉ hiện khi cần
	DurationSeconds int    `json:"duration_seconds"`
	OrderIndex      int    `json:"order_index"`
}
