package core

import "time"

// Bảng roles - Ẩn ID
type Role struct {
	ID   uint   `gorm:"primaryKey" json:"-"`
	Name string `gorm:"type:varchar(50)" json:"name"`
}

// Bảng users
type User struct {
	ID           uint      `gorm:"primaryKey" json:"-"` // Ẩn ID
	RoleID       uint      `gorm:"not null" json:"-"`   // Ẩn khóa ngoại
	Email        string    `gorm:"type:varchar(150)" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255)" json:"-"` // Luôn luôn ẩn
	FullName     string    `gorm:"type:varchar(100)" json:"full_name"`
	AvatarURL    string    `gorm:"type:varchar(255)" json:"avatar_url"`
	Phone        string    `gorm:"type:varchar(20)" json:"phone"`
	Status       string    `gorm:"type:varchar(20)" json:"status"`
	CreatedAt    time.Time `gorm:"default:current_timestamp" json:"created_at"`

	Profile *UserProfile `gorm:"foreignKey:UserID" json:"profile,omitempty"`
	Streak  *UserStreak  `gorm:"foreignKey:UserID" json:"streak,omitempty"`
}

// Bảng user_profiles - Ẩn các ID kỹ thuật
type UserProfile struct {
	ID          uint      `gorm:"primaryKey" json:"-"`
	UserID      uint      `gorm:"unique" json:"-"`
	Headline    string    `gorm:"type:varchar(150)" json:"headline"`
	Bio         string    `gorm:"type:text" json:"bio"`
	Skills      string    `gorm:"type:longtext" json:"skills"`
	GithubURL   string    `gorm:"type:varchar(255)" json:"github_url"`
	LinkedinURL string    `gorm:"type:varchar(255)" json:"linkedin_url"`
	WebsiteURL  string    `gorm:"type:varchar(255)" json:"website_url"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

// Bảng user_streaks
type UserStreak struct {
	ID            uint       `gorm:"primaryKey" json:"-"`
	UserID        uint       `gorm:"unique" json:"-"`
	CurrentStreak int        `gorm:"default:0" json:"current_streak"`
	LastStudyDate *time.Time `gorm:"type:date" json:"last_study_date"`
}

// Bảng notifications
type Notification struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	UserID    uint      `gorm:"not null" json:"-"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Type      string    `gorm:"type:varchar(50)" json:"type"`
	IsRead    bool      `gorm:"type:tinyint(1);default:0" json:"is_read"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
}
