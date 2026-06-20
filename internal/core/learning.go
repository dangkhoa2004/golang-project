package core

import "time"

// Bảng enrollments
type Enrollment struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	StudentID       uint      `gorm:"not null" json:"-"`
	CourseID        uint      `gorm:"not null" json:"-"`
	ProgressPercent int       `gorm:"default:0" json:"progress_percent"`
	Status          string    `gorm:"type:varchar(50);default:'learning'" json:"status"`
	EnrolledAt      time.Time `gorm:"default:current_timestamp" json:"enrolled_at"`

	// Preload Course để dễ dàng hiển thị danh sách khóa học của tôi
	Course *Course `gorm:"foreignKey:CourseID" json:"course,omitempty"`
}

// Bảng lesson_progress
type LessonProgress struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	EnrollmentID      uint      `gorm:"not null" json:"-"`
	LessonID          uint      `gorm:"not null" json:"-"`
	IsCompleted       bool      `gorm:"type:tinyint(1);default:0" json:"is_completed"`
	LastWatchedSecond int       `gorm:"default:0" json:"last_watched_second"`
	UpdatedAt         time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

// Bảng discussions (Q&A)
type Discussion struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	LessonID       uint      `gorm:"not null" json:"-"`
	UserID         uint      `gorm:"not null" json:"-"`
	ParentID       *uint     `json:"-"`
	Content        string    `gorm:"type:text;not null" json:"content"`
	VideoTimestamp int       `json:"video_timestamp"`
	CreatedAt      time.Time `gorm:"default:current_timestamp" json:"created_at"`
}

// Bảng user_notes
type UserNote struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	LessonID       uint      `gorm:"not null" json:"-"`
	UserID         uint      `gorm:"not null" json:"-"`
	Content        string    `gorm:"type:text;not null" json:"content"`
	VideoTimestamp int       `json:"video_timestamp"`
	CreatedAt      time.Time `gorm:"default:current_timestamp" json:"created_at"`
}

// Bảng certificates
type Certificate struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	UserID         uint      `gorm:"not null" json:"-"`
	CourseID       uint      `gorm:"not null" json:"-"`
	CertificateURL string    `gorm:"type:varchar(255)" json:"certificate_url"`
	IssuedAt       time.Time `gorm:"default:current_timestamp" json:"issued_at"`
}
