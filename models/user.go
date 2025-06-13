package models

type User struct {
	UserID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string `gorm:"type:varchar(100);not null" json:"name"`
	Contact         string `gorm:"type:varchar(15);unique;not null" json:"contact"`
	Skills          string `gorm:"type:text;not null" json:"skills"`
	Age             int    `gorm:"not null" json:"age"`
	ExperienceYears int    `gorm:"not null" json:"experience_years"`
	Education       string `gorm:"type:varchar(255);not null" json:"education"`
}
