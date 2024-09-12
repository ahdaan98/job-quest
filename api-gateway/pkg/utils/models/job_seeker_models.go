package models

type JobSeekerLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type JobSeekerSignUp struct {
	Email       string `json:"email" binding:"required" validate:"required,email"`
	Password    string `json:"password" binding:"required" validate:"min=6,max=20"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Bio         string `json:"bio"`
}

type JobSeekerDetailsResponse struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Bio         string `json:"bio"`
}

type TokenJobSeeker struct {
	JobSeeker JobSeekerDetailsResponse
	Token     string
}

type SavedJobs struct {
	JobID       int64 `json:"job_id" validate:"required"`
	JobseekerID int64 `json:"jobseeker_id" validate:"required"`
}

type SavedJobsResponse struct {
	ID          uint  `json:"id"`
	JobID       int64 `json:"job_id" validate:"required"`
	JobseekerID int64 `json:"jobseeker_id" validate:"required"`
}

type JobSeekerVerifyOTPRequest struct {
	Email string `json:"email"` // Email associated with the OTP
	Otp   int32  `json:"otp"`   // OTP to verify
}

type OtpResponse struct {
	Message string `json:"message"`
}

type OtpVerificationResponse struct {
	Status           int64                     `json:"status"`
	Token            string                    `json:"token"`
	JobSeekerDetails *JobSeekerDetailsResponse `json:"job_seeker_details"`
}

type SubscriptionPlanResponse struct {
    Success bool
    Message string
}