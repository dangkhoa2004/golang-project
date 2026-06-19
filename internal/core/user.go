package core

import "time"

// Bảng roles
type Role struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}

// Bảng users
type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleID       uint      `gorm:"not null" json:"role_id"`
	Email        string    `gorm:"type:varchar(150);unique;not null" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"` // Ẩn mật khẩu khi xuất JSON
	FullName     string    `gorm:"type:varchar(100);not null" json:"full_name"`
	AvatarURL    string    `gorm:"type:varchar(255)" json:"avatar_url"`
	Phone        string    `gorm:"type:varchar(20)" json:"phone"`
	Status       string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	CreatedAt    time.Time `gorm:"default:current_timestamp" json:"created_at"`

	Profile *UserProfile `gorm:"foreignKey:UserID" json:"profile,omitempty"`
	Streak  *UserStreak  `gorm:"foreignKey:UserID" json:"streak,omitempty"`
}

// Bảng user_profiles
type UserProfile struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"unique;not null" json:"user_id"`
	Headline    string    `gorm:"type:varchar(150)" json:"headline"`
	Bio         string    `gorm:"type:text" json:"bio"`
	Skills      string    `gorm:"type:longtext" json:"skills"` // Lưu chuỗi định dạng JSON
	GithubURL   string    `gorm:"type:varchar(255)" json:"github_url"`
	LinkedinURL string    `gorm:"type:varchar(255)" json:"linkedin_url"`
	WebsiteURL  string    `gorm:"type:varchar(255)" json:"website_url"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

// Bảng user_streaks
type UserStreak struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint       `gorm:"unique;not null" json:"user_id"`
	CurrentStreak int        `gorm:"default:0" json:"current_streak"`
	LastStudyDate *time.Time `gorm:"type:date" json:"last_study_date"` // Dùng con trỏ để cho phép NULL
}

// Bảng notifications
type Notification struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Type      string    `gorm:"type:varchar(50)" json:"type"`
	IsRead    bool      `gorm:"type:tinyint(1);default:0" json:"is_read"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
}
