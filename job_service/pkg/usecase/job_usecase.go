package usecase

import (
	"errors"
	"fmt"
	"job_service/pkg/config"
	"job_service/pkg/helper"
	interfaces "job_service/pkg/repository/interface"
	services "job_service/pkg/usecase/interface"
	"job_service/pkg/utils/models"

	"github.com/google/uuid"
)

type jobUseCase struct {
	jobRepository interfaces.JobRepository
}

func NewJobUseCase(repository interfaces.JobRepository) services.JobUseCase {
	return &jobUseCase{
		jobRepository: repository,
	}
}

func (ju *jobUseCase) PostJob(job models.JobOpening, employerID int32) (models.JobOpeningResponse, error) {

	if employerID <= 0 {
		return models.JobOpeningResponse{}, errors.New("invalid input data")
	}

	jobData, err := ju.jobRepository.PostJob(job, int32(employerID))
	if err != nil {
		return models.JobOpeningResponse{}, err
	}
	return jobData, nil
}

func (ju *jobUseCase) GetAllJobs(employerID int32) ([]models.AllJob, error) {

	if employerID <= 0 {
		return []models.AllJob{}, errors.New("invalid input data")
	}

	jobData, err := ju.jobRepository.GetAllJobs(employerID)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}

func (ju *jobUseCase) GetAJob(employerID, jobId int32) (models.JobOpeningResponse, error) {

	if employerID <= 0 || jobId <= 0 {
		return models.JobOpeningResponse{}, errors.New("invalid input data")
	}

	isJobExist, err := ju.jobRepository.IsJobExist(jobId)
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobExist {
		return models.JobOpeningResponse{}, fmt.Errorf("job with ID %d does not exist", jobId)
	}

	jobData, err := ju.jobRepository.GetAJob(employerID, jobId)
	if err != nil {
		return models.JobOpeningResponse{}, err
	}
	return jobData, nil
}

func (ju *jobUseCase) DeleteAJob(employerIDInt, jobID int32) error {

	if employerIDInt <= 0 || jobID <= 0 {
		return errors.New("invalid input data")
	}

	isJobExist, err := ju.jobRepository.IsJobExist(jobID)
	if err != nil {
		return fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobExist {
		return fmt.Errorf("job with ID %d does not exist", jobID)
	}

	err = ju.jobRepository.DeleteAJob(employerIDInt, jobID)
	if err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}

	return nil
}
func (ju *jobUseCase) UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error) {
    fmt.Printf("Updating job with Employer ID: %d, Job ID: %d\n", employerID, jobID)

    if employerID <= 0 || jobID <= 0 {
        return models.JobOpeningResponse{}, errors.New("invalid input data: employerID or jobID is less than or equal to zero")
    }

    isJobExist, err := ju.jobRepository.IsJobExist(jobID)
    if err != nil {
        return models.JobOpeningResponse{}, fmt.Errorf("failed to check if job exists: %v", err)
    }

    if !isJobExist {
        return models.JobOpeningResponse{}, fmt.Errorf("job with ID %d does not exist", jobID)
    }

    updatedJob, err := ju.jobRepository.UpdateAJob(employerID, jobID, jobDetails)
    if err != nil {
        return models.JobOpeningResponse{}, fmt.Errorf("failed to update job: %v", err)
    }

    return updatedJob, nil
}


func (ju *jobUseCase) JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error) {

	// if keyword == "" {
	// 	return []models.JobSeekerGetAllJobs{}, errors.New("invalid input data")
	// }

	jobs, err := ju.jobRepository.JobSeekerGetAllJobs(keyword)
	if err != nil {
		return nil, fmt.Errorf("failed to get jobs: %v", err)
	}

	var jobSeekerJobs []models.JobSeekerGetAllJobs
	for _, job := range jobs {

		jobSeekerJob := models.JobSeekerGetAllJobs{
			ID:    job.ID,
			Title: job.Title,
		}
		jobSeekerJobs = append(jobSeekerJobs, jobSeekerJob)
	}

	return jobSeekerJobs, nil
}

func (ju *jobUseCase) GetJobDetails(jobID int32) (models.JobOpeningResponse, error) {

	if jobID <= 0 {
		return models.JobOpeningResponse{}, errors.New("invalid input data")
	}

	isJobExist, err := ju.jobRepository.IsJobExist(jobID)
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobExist {
		return models.JobOpeningResponse{}, fmt.Errorf("job with ID %d does not exist", jobID)
	}

	jobData, err := ju.jobRepository.GetJobDetails(jobID)
	if err != nil {
		return models.JobOpeningResponse{}, err
	}

	return jobData, nil
}

func (ju *jobUseCase) ApplyJob(jobApplication models.ApplyJob, resumeData []byte) (models.ApplyJobResponse, error) {

	if jobApplication.JobID <= 0 || jobApplication.JobseekerID <= 0 || jobApplication.CoverLetter == "" {
		return models.ApplyJobResponse{}, errors.New("invalid input data")
	}

	fileUID := uuid.New()
	fileName := fileUID.String()
	h := helper.NewHelper(config.Config{})

	url, err := h.AddImageToAwsS3(resumeData, fileName)
	if err != nil {
		return models.ApplyJobResponse{}, err
	}

	fmt.Println("url", url)

	Data, err := ju.jobRepository.ApplyJob(jobApplication, url)
	if err != nil {
		return models.ApplyJobResponse{}, err
	}

	return Data, nil
}

func (ju *jobUseCase) GetApplicants(employerID int64) ([]models.ApplyJobResponse, error) {

	if employerID <= 0 {
		return []models.ApplyJobResponse{}, errors.New("cannot use negative values")
	}

	jobid, err := ju.jobRepository.GetJobIDByEmployerID(employerID)
	if err != nil {
		return nil, fmt.Errorf("failed to check if job exists: %v", err)
	}

	fmt.Println("jobid", jobid)

	var allApplicants []models.ApplyJobResponse
    for _, jobID := range jobid {
        applicants, err := ju.jobRepository.GetApplicantsByEmployerID(int64(jobID.ID))
        if err != nil {
            return nil, fmt.Errorf("failed to get applicants for job ID %d: %v", jobID.ID, err)
        }
        allApplicants = append(allApplicants, applicants...)
    }

    return allApplicants, nil
}

func (uc *jobUseCase) SaveJobs(jobID, userID int64) (models.SavedJobsResponse, error) {

	if jobID <= 0 || userID <= 0 {
		return models.SavedJobsResponse{}, errors.New("cannot use negative values")
	}

	isJobAvailable, err := uc.jobRepository.IsJobExist(int32(jobID))
	if err != nil {
		return models.SavedJobsResponse{}, fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobAvailable {
		return models.SavedJobsResponse{}, errors.New("job does not exist")
	}

	savedJob, err := uc.jobRepository.SaveJobs(jobID, userID)
	if err != nil {
		return models.SavedJobsResponse{}, fmt.Errorf("failed to save job: %v", err)
	}

	response := models.SavedJobsResponse{
		ID:          uint(savedJob.ID),
		JobID:       savedJob.JobID,
		JobseekerID: savedJob.JobseekerID,
	}
	return response, nil
}

func (ju *jobUseCase) DeleteSavedJob(jobID, userID int32) error {

	if jobID <= 0 || userID <= 0 {
		return errors.New("cannot use negative values")
	}

	isJobSaved, err := ju.jobRepository.IsJobSaved(jobID, userID)
	if err != nil {
		return fmt.Errorf("failed to check if job is saved: %v", err)
	}

	if !isJobSaved {
		return errors.New("job is not saved by the user")
	}

	err = ju.jobRepository.DeleteSavedJob(jobID, userID)
	if err != nil {
		return fmt.Errorf("failed to delete saved job: %v", err)
	}
	return nil
}

func (ju *jobUseCase) GetSavedJobs(userIdInt int32) ([]models.SavedJobsResponse, error) {

	if userIdInt <= 0 {
		return []models.SavedJobsResponse{}, errors.New("cannot use negative values")
	}

	savedJobs, err := ju.jobRepository.GetSavedJobs(userIdInt)
	if err != nil {
		return nil, fmt.Errorf("failed to get saved jobs: %v", err)
	}

	return savedJobs, nil
}

func (ju *jobUseCase) UpdateApplyJob(applyJobID uint, status string) (uint,uint,error) {
	if applyJobID <= 0 {
		return 0,0, errors.New("invalid applyJobID")
	}

	// Update the status of the job application and retrieve the jobSeekerID
	jobSeekerID,jobID, err := ju.jobRepository.UpdateApplyJobStatus(applyJobID, status)
	if err != nil {
		return 0,0, fmt.Errorf("failed to update job application status: %v", err)
	}

	// Return the jobSeekerID along with a nil error if successful
	return jobSeekerID,jobID, nil
}

func (ju *jobUseCase) GetApplicantsByStatus(jobID int64, status string) ([]models.ApplyJobResponse, error) {
	if jobID <= 0 {
		return nil, errors.New("invalid jobID")
	}

	// Retrieve the list of accepted applicants for the given jobID
	acceptedApplicants, err := ju.jobRepository.GetApplicants(jobID,status)
	if err != nil {
		return nil, fmt.Errorf("failed to get accepted applicants: %v", err)
	}

	return acceptedApplicants, nil
}