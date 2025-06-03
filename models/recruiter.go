package models

type Recruiter struct {
	RecruiterID uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"size:100;not null"`
	Contact     string `gorm:"size:20;not null"`
	CompanyName string `gorm:"size:255;not null"`
}
