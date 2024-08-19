package domain

import "time"

type JobOpening struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title" gorm:"type:varchar(255);not null"`
	Description         string    `json:"description" gorm:"type:text;not null"`
	Requirements        string    `json:"requirements" gorm:"type:text;not null"`
	PostedOn            time.Time `json:"posted_on" gorm:"not null"`
	EmployerID          int32     `json:"employer_id" gorm:"not null"`
	Location            string    `json:"location" gorm:"type:varchar(255);not null"`
	EmploymentType      string    `json:"employment_type" gorm:"type:varchar(100);not null"`
	Salary              string    `json:"salary" gorm:"type:varchar(100);not null"`
	SkillsRequired      string    `json:"skills_required" gorm:"-"`
	ExperienceLevel     string    `json:"experience_level" gorm:"type:varchar(100);not null"`
	EducationLevel      string    `json:"education_level" gorm:"type:varchar(100);not null"`
	ApplicationDeadline time.Time `json:"application_deadline" gorm:"not null"`
}

type JobOpeningResponse struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	PostedOn            time.Time `json:"posted_on"`
	EmployerID          int       `json:"employer_id"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	Salary              string    `json:"salary"`
	SkillsRequired      string    `json:"skills_required"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
	UpdatedOn           time.Time `json:"updated_on"`
	IsDeleted           bool      `json:"is_deleted"`
}