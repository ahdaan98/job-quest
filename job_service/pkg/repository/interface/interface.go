package interfaces

import "job_service/pkg/utils/models"

type JobRepository interface {
	PostJob(jobDetails models.JobOpening, employerID int32) (models.JobOpeningResponse, error)
	GetAllJobs(employerID int32) ([]models.AllJob, error)
	GetAJob(employerID, jobId int32) (models.JobOpeningResponse, error)
	IsJobExist(jobID int32) (bool, error)
	GetApplicantsByEmployerID(employerID int64) ([]models.ApplyJobResponse, error)
	GetJobIDByEmployerID(employerID int64) ([]models.JobOpeningResponse, error)
	DeleteAJob(employerIDInt, jobID int32) error
	JobSeekerGetAllJobs(keyword string) ([]models.JobOpeningResponse, error)
	GetJobDetails(jobID int32) (models.JobOpeningResponse, error)
	ApplyJob(application models.ApplyJob, resumeURL string) (models.ApplyJobResponse, error)
	SaveJobs(jobID, userID int64) (models.SavedJobsResponse, error)
	IsJobSaved(jobID, userID int32) (bool, error)
	DeleteSavedJob(jobID, userID int32) error
	GetSavedJobs(userIdInt int32) ([]models.SavedJobsResponse, error)
	UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error)
	UpdateApplyJobStatus(applyJobID uint, status string) (uint, uint, error)
	GetApplicants(JobID int64, status string) ([]models.ApplyJobResponse,error)
}