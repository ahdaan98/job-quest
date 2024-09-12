package interfaces

import (
	"mime/multipart"

	"github.com/ahdaan67/JobQuest/pkg/utils/models"
)

type JobClient interface {
	PostJobOpening(jobDetails models.JobOpening, EmployerID int32) (models.JobOpeningResponse, error)
	GetAllJobs(employerIDInt int32) ([]models.AllJob, error)
	GetAJob(employerIDInt, jobId int32) (models.JobOpeningResponse, error)
	DeleteAJob(employerIDInt, jobID int32) error
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
	GetJobDetails(jobID int32) (models.JobOpeningResponse, error)
	UpdateAJob(employerIDInt int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error)
	ApplyJob(jobApplication models.ApplyJob, file *multipart.FileHeader) (models.ApplyJobResponse, error)
	GetApplicants(employerID int64) ([]models.ApplyJobResponse, error)
	SaveAJob(userIdInt, jobIdInt int32) (models.SavedJobsResponse, error)
	DeleteSavedJob(jobIdInt, userIdInt int32) error
	GetASavedJob(userIdInt int32) ([]models.SavedJobsResponse, error)
	UpdateApplyJob(applyJobID uint64, status string) (uint, uint, error)
	GetAcceptedApplicants(jobID int64, status string) ([]models.ApplyJobResponse, error)
}
