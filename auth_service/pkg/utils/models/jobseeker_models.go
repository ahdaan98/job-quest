package models

import "time"

type JobSeekerLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type JobSeekerSignUp struct {
	ID          uint   `json:"id"`
	Email       string `json:"email" binding:"required" validate:"required,email"`
	Password    string `json:"password" binding:"required" validate:"min=6,max=20"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
}

type JobSeekerDetailsResponse struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
}

type TokenJobSeeker struct {
	JobSeeker JobSeekerDetailsResponse
	Token     string
}

type JobseekerSubscription struct {
	ID                       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	JobSeekerID              uint      `gorm:"not null;index" json:"jobseekerId"`
	SubscriptionPlanID       uint      `gorm:"not null;index" json:"subscriptionPlanId"`
	RemainingJobApplications int       `gorm:"not null" json:"remainingJobApplications"`
	StartDate                time.Time `gorm:"not null;type:timestamp" json:"startDate"`
	EndDate                  time.Time `gorm:"not null;type:timestamp" json:"endDate"`
	IsActive                 bool      `gorm:"not null;default:true" json:"isActive"`
}

type JobseekerPlan struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title          string `gorm:"not null" json:"title"`
	ApplyJobs      int    `gorm:"not null" json:"apply_jobs"`
	AdvancedFilter bool   `gorm:"not null;default:false" json:"advanced_filter"`
}