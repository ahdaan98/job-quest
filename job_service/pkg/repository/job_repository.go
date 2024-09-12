package repository

import (
	"errors"
	"fmt"
	"job_service/pkg/domain"
	interfaces "job_service/pkg/repository/interface"
	"job_service/pkg/utils/models"
	"time"

	"gorm.io/gorm"
)

type jobRepository struct {
	DB *gorm.DB
}

func NewJobRepository(DB *gorm.DB) interfaces.JobRepository {
	return &jobRepository{
		DB: DB,
	}
}

func (jr *jobRepository) PostJob(jobDetails models.JobOpening, employerID int32) (models.JobOpeningResponse, error) {
	postedOn := time.Now()

	job := models.JobOpeningResponse{
		Title:               jobDetails.Title,
		Description:         jobDetails.Description,
		Requirements:        jobDetails.Requirements,
		PostedOn:            postedOn,
		EmployerID:          int(employerID),
		Location:            jobDetails.Location,
		EmploymentType:      jobDetails.EmploymentType,
		Salary:              jobDetails.Salary,
		SkillsRequired:      jobDetails.SkillsRequired,
		ExperienceLevel:     jobDetails.ExperienceLevel,
		EducationLevel:      jobDetails.EducationLevel,
		ApplicationDeadline: jobDetails.ApplicationDeadline,
	}

	if err := jr.DB.Create(&job).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return job, nil
}

func (jr *jobRepository) GetAllJobs(employerID int32) ([]models.AllJob, error) {
	var jobs []models.AllJob

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("employer_id = ?", employerID).Select("id, title, application_deadline, employer_id").Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (jr *jobRepository) GetAJob(employerID, jobId int32) (models.JobOpeningResponse, error) {
	var job models.JobOpeningResponse

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ? AND employer_id = ?", jobId, employerID).First(&job).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return job, nil
}

func (jr *jobRepository) IsJobExist(jobID int32) (bool, error) {
	var job models.JobOpeningResponse

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ?", jobID).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (jr *jobRepository) GetJobIDByEmployerID(employerID int64) ([]models.JobOpeningResponse, error) {
	var job []models.JobOpeningResponse

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("employer_id = ?", employerID).Scan(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return job, nil
}

func (jr *jobRepository) GetApplicantsByEmployerID(jobID int64) ([]models.ApplyJobResponse, error) {
	var applicants []models.ApplyJobResponse

	fmt.Println("Job ID:", jobID)

	query := "SELECT * FROM apply_jobs WHERE job_id = ?"
	if err := jr.DB.Raw(query, jobID).Scan(&applicants).Error; err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	// Check if applicants are retrieved
	if len(applicants) == 0 {
		fmt.Println("No applicants found for job ID:", jobID)
	} else {
		fmt.Println("Retrieved applicants:", applicants)
	}

	return applicants, nil
}

func (jr *jobRepository) DeleteAJob(employerID, jobID int32) error {
	if err := jr.DB.Where("id = ? AND employer_id = ?", jobID, employerID).Delete(&models.JobOpeningResponse{}).Error; err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}

	return nil
}

func (jr *jobRepository) JobSeekerGetAllJobs(keyword string) ([]models.JobOpeningResponse, error) {
	var jobSeekerJobs []models.JobOpeningResponse

	if err := jr.DB.Find(&jobSeekerJobs).Error; err != nil {
		return nil, fmt.Errorf("failed to query jobs: %v", err)
	}

	return jobSeekerJobs, nil
}

func (jr *jobRepository) GetJobDetails(jobID int32) (models.JobOpeningResponse, error) {
	var job models.JobOpeningResponse

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ?", jobID).First(&job).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return job, nil
}

func (jr *jobRepository) ApplyJob(application models.ApplyJob, resumeURL string) (models.ApplyJobResponse, error) {
	var jobResponse models.ApplyJobResponse

	// Debug: Print the input values
	fmt.Printf("Applying for job with jobseeker_id=%d and job_id=%d\n", application.JobseekerID, application.JobID)

	// Check if the job has already been applied
	var count int64
	err := jr.DB.Model(&models.ApplyJob{}).
		Where("jobseeker_id = ? AND job_id = ?", application.JobseekerID, application.JobID).
		Count(&count).Error
	if err != nil {
		return models.ApplyJobResponse{}, fmt.Errorf("error checking if job is already applied: %w", err)
	}

	// Debug: Print the count of existing applications
	fmt.Printf("Count of existing applications: %d\n", count)

	if count > 0 {
		return models.ApplyJobResponse{}, errors.New("job already applied")
	}

	// Insert new job application
	result := jr.DB.Exec("INSERT INTO apply_jobs (jobseeker_id, job_id, resume_url, cover_letter) VALUES (?, ?, ?, ?)",
		application.JobseekerID,
		application.JobID,
		resumeURL,
		application.CoverLetter)

	if result.Error != nil {
		return models.ApplyJobResponse{}, fmt.Errorf("error inserting into database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return models.ApplyJobResponse{}, errors.New("no rows were affected during insert")
	}

	// Retrieve the inserted record
	err = jr.DB.Raw("SELECT id, jobseeker_id, job_id, resume_url, cover_letter FROM apply_jobs WHERE jobseeker_id = ? AND job_id = ?",
		application.JobseekerID,
		application.JobID).Scan(&jobResponse).Error
	if err != nil {
		return models.ApplyJobResponse{}, fmt.Errorf("failed to retrieve inserted record: %w", err)
	}

	// Debug: Print the retrieved job response
	fmt.Printf("Retrieved job response: %+v\n", jobResponse)

	return jobResponse, nil
}

func (jr *jobRepository) SaveJobs(jobID, userID int64) (models.SavedJobsResponse, error) {

	var savedJobResponse models.SavedJobsResponse

	result := jr.DB.Exec("INSERT INTO saved_jobs (job_id, jobseeker_id) VALUES (?, ?) ", jobID, userID)
	if result.Error != nil {
		return models.SavedJobsResponse{}, fmt.Errorf("error inserting into database: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return models.SavedJobsResponse{}, errors.New("no rows were affected during insert")
	}

	err := jr.DB.Raw("SELECT * FROM saved_jobs WHERE job_id = ? AND jobseeker_id = ?", jobID, userID).Scan(&savedJobResponse).Error
	if err != nil {
		return models.SavedJobsResponse{}, fmt.Errorf("failed to retrieve saved job: %w", err)
	}

	return savedJobResponse, nil
}

func (jr *jobRepository) IsJobSaved(jobID, userID int32) (bool, error) {
	var savedJob models.SavedJobsResponse
	err := jr.DB.Raw("SELECT * FROM saved_jobs WHERE job_id = ? AND jobseeker_id = ?", jobID, userID).Scan(&savedJob).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("failed to retrieve saved job: %w", err)
	}
	return true, nil
}

func (jr *jobRepository) DeleteSavedJob(jobID, userID int32) error {
	result := jr.DB.Exec("DELETE FROM saved_jobs WHERE job_id = ? AND jobseeker_id = ?", jobID, userID)
	if result.Error != nil {
		return fmt.Errorf("error deleting saved job: %w", result.Error)
	}
	return nil
}

func (jr *jobRepository) GetSavedJobs(userID int32) ([]models.SavedJobsResponse, error) {
	var savedJobs []models.SavedJobsResponse
	err := jr.DB.Raw("SELECT * FROM saved_jobs WHERE jobseeker_id = ?", userID).Scan(&savedJobs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve saved jobs: %w", err)
	}
	return savedJobs, nil
}

func (jr *jobRepository) UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error) {
	postedOn := time.Now()

	updatedJob := models.JobOpeningResponse{
		ID:                  uint(jobID),
		Title:               jobDetails.Title,
		Description:         jobDetails.Description,
		Requirements:        jobDetails.Requirements,
		PostedOn:            postedOn,
		EmployerID:          int(employerID),
		Location:            jobDetails.Location,
		EmploymentType:      jobDetails.EmploymentType,
		Salary:              jobDetails.Salary,
		SkillsRequired:      jobDetails.SkillsRequired,
		ExperienceLevel:     jobDetails.ExperienceLevel,
		EducationLevel:      jobDetails.EducationLevel,
		ApplicationDeadline: jobDetails.ApplicationDeadline,
	}

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ? AND employer_id = ?", jobID, employerID).Updates(updatedJob).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return updatedJob, nil
}

func (jr *jobRepository) UpdateApplyJobStatus(applyJobID uint, status string) (uint, uint, error) {
	var jobSeekerID uint
	var jobID uint

	// Fetch the jobSeekerID and jobID based on applyJobID
	if err := jr.DB.Model(&domain.ApplyJob{}).
		Select("jobseeker_id").
		Where("id = ?", applyJobID).
		Scan(&jobSeekerID).Error; err != nil {
		return 0, 0, err
	}

	if err := jr.DB.Model(&domain.ApplyJob{}).
		Select("job_id").
		Where("id = ?", applyJobID).
		Scan(&jobID).Error; err != nil {
		return 0, 0, err
	}

	// Update the status of the job application
	if err := jr.DB.Model(&domain.ApplyJob{}).Where("id = ?", applyJobID).Update("status", status).Error; err != nil {
		return 0, 0, err
	}

	return jobSeekerID, jobID, nil
}
func (jr *jobRepository) GetApplicants(JobID int64, status string) ([]models.ApplyJobResponse, error) {
	var acceptedApplicants []models.ApplyJobResponse

	// Define the query to select the accepted applicants for the given jobID
	query := `SELECT id, jobseeker_id, job_id, resume_url, cover_letter 
              FROM apply_jobs 
              WHERE job_id = ? AND status = ?`

	// Execute the query and scan the results into the acceptedApplicants slice
	err := jr.DB.Raw(query, JobID, status).Scan(&acceptedApplicants).Error
	if err != nil {
		return nil, err
	}

	return acceptedApplicants, nil
}
