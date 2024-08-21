package interfaces

import "github.com/ahdaan67/JobQuest/pkg/utils/models"

type JobSeekerClient interface {
	JobSeekerSignUp(jobSeekerDetails models.JobSeekerSignUp) (models.TokenJobSeeker, error)
	JobSeekerLogin(jobSeekerDetails models.JobSeekerLogin) (models.TokenJobSeeker, error)
}