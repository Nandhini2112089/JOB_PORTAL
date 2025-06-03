package models

import "time"

type Interview struct {
	InterviewID   uint      `gorm:"primaryKey;autoIncrement"`
	ApplicationID uint      `gorm:"not null"`
	UserID        uint      `gorm:"not null"`
	JobID         uint      `gorm:"not null"`
	CompanyName   string    `gorm:"not null"`
	RecruiterID   uint      `gorm:"not null"`
	RecruiterName string    `gorm:"not null"`
	InterviewTime time.Time `gorm:"not null"`
}
