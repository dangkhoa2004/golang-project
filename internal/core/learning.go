package core

import "time"

// Bảng enrollments
type Enrollment struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID       uint      `gorm:"not null" json:"student_id"`
	CourseID        uint      `gorm:"not null" json:"course_id"`
	ProgressPercent int       `gorm:"default:0" json:"progress_percent"`
	Status          string    `gorm:"type:varchar(50);default:'learning'" json:"status"`
	EnrolledAt      time.Time `gorm:"default:current_timestamp" json:"enrolled_at"`
}

// Bảng lesson_progress
type LessonProgress struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	EnrollmentID      uint      `gorm:"not null" json:"enrollment_id"`
	LessonID          uint      `gorm:"not null" json:"lesson_id"`
	IsCompleted       bool      `gorm:"type:tinyint(1);default:0" json:"is_completed"`
	LastWatchedSecond int       `gorm:"default:0" json:"last_watched_second"`
	UpdatedAt         time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

// Bảng discussions
type Discussion struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	LessonID       uint      `gorm:"not null" json:"lesson_id"`
	UserID         uint      `gorm:"not null" json:"user_id"`
	ParentID       *uint     `json:"parent_id"` // Dùng con trỏ đại diện cho khóa ngoại đệ quy có thể NULL
	Content        string    `gorm:"type:text;not null" json:"content"`
	VideoTimestamp int       `json:"video_timestamp"`
	CreatedAt      time.Time `gorm:"default:current_timestamp" json:"created_at"`
}

// Bảng user_notes
type UserNote struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	LessonID       uint      `gorm:"not null" json:"lesson_id"`
	UserID         uint      `gorm:"not null" json:"user_id"`
	Content        string    `gorm:"type:text;not null" json:"content"`
	VideoTimestamp int       `json:"video_timestamp"`
	CreatedAt      time.Time `gorm:"default:current_timestamp" json:"created_at"`
}

// Bảng certificates
type Certificate struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         uint      `gorm:"not null" json:"user_id"`
	CourseID       uint      `gorm:"not null" json:"course_id"`
	CertificateURL string    `gorm:"type:varchar(255)" json:"certificate_url"`
	IssuedAt       time.Time `gorm:"default:current_timestamp" json:"issued_at"`
}
