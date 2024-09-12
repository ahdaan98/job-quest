package interfaces

import "github.com/ahdaan67/JobQuest/pkg/utils/models"

type JobSeekerClient interface {
	JobSeekerSignUp(jobSeekerDetails models.JobSeekerSignUp) (models.TokenJobSeeker, error)
	JobSeekerLogin(jobSeekerDetails models.JobSeekerLogin) (models.TokenJobSeeker, error)
	JobSeekerOTPSignUp(jobSeekerDetails models.JobSeekerSignUp) (models.OtpResponse, error) // Added
	JobSeekerVerifyOTP(otpRequest models.JobSeekerVerifyOTPRequest) (models.OtpVerificationResponse, error) // Added
	JobSeekerLinkedinSign(jobSeekerDetails models.JobSeekerDetailsResponse) (models.TokenJobSeeker, error)
	GetEmailByJobSeekerID(id uint) (string, error)
	ActivateSubscriptionPlan(jobSeekerID uint, subscriptionPlanID uint) (models.SubscriptionPlanResponse, error)
}