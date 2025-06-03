package models

import "time"

type Application struct {
	ApplicationID uint      `gorm:"primaryKey;autoIncrement"`
	UserID        uint      `gorm:"not null"`
	JobID         uint      `gorm:"not null"`
	Status        string    `gorm:"type:enum('Pending', 'Shortlisted', 'Rejected');default:'Pending';not null"`
	AppliedAt     time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}
