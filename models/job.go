package models

type Job struct {
	JobID          uint    `gorm:"primaryKey;autoIncrement"`
	RecruiterID    uint    `gorm:"not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;foreignKey:RecruiterID;references:ID"`
	Title          string  `gorm:"size:255;not null"`
	Description    string  `gorm:"type:text;not null"`
	SkillsRequired string  `gorm:"type:text;not null"`
	Location       string  `gorm:"size:255;not null"`
	Salary         float64 `gorm:"not null"`
	JobType        string  `gorm:"type:enum('Full-time', 'Part-time', 'Contract', 'Internship');not null"`
}
