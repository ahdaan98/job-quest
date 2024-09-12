package domain

import (
	"Auth/pkg/utils/models"
	"time"
)

type Admin struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
}
type TokenAdmin struct {
	Admin models.AdminDetailsResponse
	Token string
}

type Employer struct {
	ID                   uint   `json:"id" gorm:"uniquekey; not null"`
	Company_name         string `json:"company_name" gorm:"validate:required"`
	Industry             string `json:"industry" gorm:"validate:required"`
	Company_size         int    `json:"company_size" gorm:"validate:required"`
	Website              string `json:"website"`
	Headquarters_address string `json:"headquarters_address"`
	About_company        string `json:"about_company" gorm:"type:text"`
	Contact_email        string `json:"contact_email" gorm:"validate:required"`
	Contact_phone_number uint   `json:"contact_phone_number" gorm:"type:numeric"`
	Password             string `json:"password" gorm:"validate:required"`
}

type TokenEmployer struct {
	Employer models.EmployerDetailsResponse
	Token    string
}

type JobSeeker struct {
	ID            uint   `json:"id" gorm:"uniquekey; not null"`
	Email         string `json:"email" gorm:"validate:required"`
	Password      string `json:"password" gorm:"validate:required"`
	First_name    string `json:"first_name" gorm:"validate:required"`
	Last_name     string `json:"last_name" gorm:"validate:required"`
	Phone_number  string `json:"phone_number" gorm:"validate:required"`
	Date_of_birth string `json:"date_of_birth" gorm:"validate:required"`
	Gender        string `json:"gender" gorm:"validate:required"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	Deleted_at    string `json:"deleted_at"`
}

type TokenJobSeeker struct {
	JobSeeker models.JobSeekerDetailsResponse
	Token     string
}

type Connections struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint
	FriendID uint
	Status   string `gorm:"check:status IN ('pending', 'blocked')"`
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
