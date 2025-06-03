package models

type User struct {
	UserID          int    `gorm:"primaryKey;autoIncrement"json:"id"`
	Name            string `gorm:"type:varchar(100);not null"`
	Contact         string `gorm:"type:varchar(15);unique;not null"`
	Skills          string `gorm:"type:text;not null"`
	Age             int    `gorm:"not null"`
	ExperienceYears int    `gorm:"not null"`
	Education       string `gorm:"type:varchar(255);not null"`
}
