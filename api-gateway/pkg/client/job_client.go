package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strconv"

	interfaces "github.com/ahdaan67/JobQuest/pkg/client/interface"
	"github.com/ahdaan67/JobQuest/pkg/config"
	pb "github.com/ahdaan67/JobQuest/pkg/pb/job"
	"github.com/ahdaan67/JobQuest/pkg/utils/models"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type jobClient struct {
	Client pb.JobClient
}

func NewJobClient(cfg config.Config) interfaces.JobClient {
	grpcConnection, err := grpc.Dial(cfg.JobQuestJob, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewJobClient(grpcConnection)

	return &jobClient{
		Client: grpcClient,
	}
}
func (jc *jobClient) PostJobOpening(jobDetails models.JobOpening, EmployerID int32) (models.JobOpeningResponse, error) {

	applicationDeadline := timestamppb.New(jobDetails.ApplicationDeadline)

	job, err := jc.Client.PostJob(context.Background(), &pb.JobOpeningRequest{
		Title:               jobDetails.Title,
		Description:         jobDetails.Description,
		Requirements:        jobDetails.Requirements,
		Location:            jobDetails.Location,
		EmploymentType:      jobDetails.EmploymentType,
		Salary:              jobDetails.Salary,
		SkillsRequired:      jobDetails.SkillsRequired,
		ExperienceLevel:     jobDetails.ExperienceLevel,
		EducationLevel:      jobDetails.EducationLevel,
		ApplicationDeadline: applicationDeadline,
		EmployerId:          EmployerID,
	})
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to post job opening: %v", err)
	}

	postedOnTime := job.PostedOn.AsTime()
	applicationDeadlineTime := job.ApplicationDeadline.AsTime()

	return models.JobOpeningResponse{
		ID:                  uint(job.Id),
		Title:               job.Title,
		Description:         job.Description,
		Requirements:        job.Requirements,
		PostedOn:            postedOnTime,
		Location:            job.Location,
		EmploymentType:      job.EmploymentType,
		Salary:              job.Salary,
		SkillsRequired:      job.SkillsRequired,
		ExperienceLevel:     job.ExperienceLevel,
		EducationLevel:      job.EducationLevel,
		ApplicationDeadline: applicationDeadlineTime,
		EmployerID:          EmployerID, // Uncomment this line if you need to set EmployerID
	}, nil
}

func (jc *jobClient) GetAllJobs(employerIDInt int32) ([]models.AllJob, error) {

	resp, err := jc.Client.GetAllJobs(context.Background(), &pb.GetAllJobsRequest{EmployerIDInt: employerIDInt})
	if err != nil {
		return nil, fmt.Errorf("failed to get all jobs: %v", err)
	}

	var allJobs []models.AllJob
	for _, job := range resp.Jobs {

		applicationDeadlineTime := job.ApplicationDeadline.AsTime()

		allJobs = append(allJobs, models.AllJob{
			ID:                  uint(job.Id),
			Title:               job.Title,
			ApplicationDeadline: applicationDeadlineTime,
			EmployerID:          employerIDInt,
		})
	}

	return allJobs, nil
}

func (jc *jobClient) GetAJob(employerIDInt, jobId int32) (models.JobOpeningResponse, error) {
	resp, err := jc.Client.GetAJob(context.Background(), &pb.GetAJobRequest{EmployerIDInt: employerIDInt, JobId: jobId})
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to get job: %v", err)
	}

	postedOnTime := resp.PostedOn.AsTime()
	applicationDeadlineTime := resp.ApplicationDeadline.AsTime()

	return models.JobOpeningResponse{
		ID:                  uint(resp.Id),
		Title:               resp.Title,
		Description:         resp.Description,
		Requirements:        resp.Requirements,
		PostedOn:            postedOnTime,
		Location:            resp.Location,
		EmploymentType:      resp.EmploymentType,
		Salary:              resp.Salary,
		SkillsRequired:      resp.SkillsRequired,
		ExperienceLevel:     resp.ExperienceLevel,
		EducationLevel:      resp.EducationLevel,
		ApplicationDeadline: applicationDeadlineTime,
		EmployerID:          employerIDInt,
	}, nil
}

func (jc *jobClient) DeleteAJob(employerIDInt, jobID int32) error {
	_, err := jc.Client.DeleteAJob(context.Background(), &pb.DeleteAJobRequest{EmployerIDInt: employerIDInt, JobId: jobID})
	if err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}
	return nil
}

func (jc *jobClient) UpdateAJob(employerIDInt int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error) {

	applicationDeadline := timestamppb.New(jobDetails.ApplicationDeadline)
	fmt.Println("emp", employerIDInt)

	job, err := jc.Client.UpdateAJob(context.Background(), &pb.UpdateAJobRequest{
		Title:               jobDetails.Title,
		Description:         jobDetails.Description,
		Requirements:        jobDetails.Requirements,
		Location:            jobDetails.Location,
		EmploymentType:      jobDetails.EmploymentType,
		Salary:              jobDetails.Salary,
		SkillsRequired:      jobDetails.SkillsRequired,
		ExperienceLevel:     jobDetails.ExperienceLevel,
		EducationLevel:      jobDetails.EducationLevel,
		ApplicationDeadline: applicationDeadline,
		EmployerId:          employerIDInt,
		JobId:               jobID,
	})
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to post job opening: %v", err)
	}

	postedOnTime := job.PostedOn.AsTime()
	applicationDeadlineTime := job.ApplicationDeadline.AsTime()

	return models.JobOpeningResponse{
		ID:                  uint(job.Id),
		Title:               job.Title,
		Description:         job.Description,
		Requirements:        job.Requirements,
		PostedOn:            postedOnTime,
		Location:            job.Location,
		EmploymentType:      job.EmploymentType,
		Salary:              job.Salary,
		SkillsRequired:      job.SkillsRequired,
		ExperienceLevel:     job.ExperienceLevel,
		EducationLevel:      job.EducationLevel,
		ApplicationDeadline: applicationDeadlineTime,
		EmployerID:          employerIDInt,
	}, nil

}

func (jc *jobClient) JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error) {
	resp, err := jc.Client.JobSeekerGetAllJobs(context.Background(), &pb.JobSeekerGetAllJobsRequest{
		Title: keyword,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get job: %v", err)
	}

	var jobs []models.JobSeekerGetAllJobs
	for _, job := range resp.Jobs {
		jobs = append(jobs, models.JobSeekerGetAllJobs{
			ID:    uint(job.Id),
			Title: job.Title,
		})
	}

	return jobs, nil
}

func (jc *jobClient) GetJobDetails(jobID int32) (models.JobOpeningResponse, error) {
	resp, err := jc.Client.GetJobDetails(context.Background(), &pb.GetJobDetailsRequest{JobId: jobID})
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to get job details: %v", err)
	}

	applicationDeadlineTime := resp.ApplicationDeadline.AsTime()

	return models.JobOpeningResponse{
		ID:                  uint(jobID),
		Title:               resp.Title,
		Description:         resp.Description,
		Requirements:        resp.Requirements,
		Location:            resp.Location,
		EmploymentType:      resp.EmploymentType,
		Salary:              resp.Salary,
		SkillsRequired:      resp.SkillsRequired,
		ExperienceLevel:     resp.ExperienceLevel,
		EducationLevel:      resp.EducationLevel,
		ApplicationDeadline: applicationDeadlineTime,
		EmployerID:          resp.EmployerId,
	}, nil
}

func (jc *jobClient) ApplyJob(jobApplication models.ApplyJob, file *multipart.FileHeader) (models.ApplyJobResponse, error) {
	var response models.ApplyJobResponse

	f, err := file.Open()
	if err != nil {
		return response, err
	}
	defer f.Close()

	fileData, err := io.ReadAll(f)
	if err != nil {
		return response, err
	}

	req := &pb.ApplyJobRequest{
		JobId:       jobApplication.JobID,
		JobseekerId: jobApplication.JobseekerID,
		CoverLetter: jobApplication.CoverLetter,
		ResumeData:  fileData,
	}

	fmt.Println("job id:",req.JobId)

	grpcResponse, err := jc.Client.ApplyJob(context.Background(), req)
	if err != nil {
		return response, err
	}

	response = models.ApplyJobResponse{
		ID:          uint(grpcResponse.Id),
		JobID:       grpcResponse.JobId,
		JobseekerID: grpcResponse.JobseekerId,
		CoverLetter: grpcResponse.CoverLetter,
		ResumeURL:   grpcResponse.ResumeUrl,
	}

	return response, nil
}

func (jc *jobClient) GetApplicants(employerID int64) ([]models.ApplyJobResponse, error) {
	var response []models.ApplyJobResponse
	req := &pb.GetJobApplicationsRequest{
		EmployerId: strconv.FormatInt(employerID, 10),
	}
	grpcResponse, err := jc.Client.GetJobApplications(context.Background(), req)
	if err != nil {
		return response, err
	}
	for _, v := range grpcResponse.JobApplications {
		jobID, err := strconv.ParseUint(v.JobId, 10, 64)
		if err != nil {
			return response, err
		}
		jobseekerID, err := strconv.ParseUint(v.JobSeekerId, 10, 64)
		if err != nil {
			return response, err
		}
		applicationID, err := strconv.ParseUint(v.Id, 10, 64)
		if err != nil {
			return response, err
		}
		response = append(response, models.ApplyJobResponse{
			ID:          uint(applicationID),
			JobID:       int64(jobID),
			JobseekerID: int64(jobseekerID),
			CoverLetter: v.CoverLetter,
			ResumeURL:   v.Resume,
		})
	}
	return response, nil
}

func (jc *jobClient) SaveAJob(userIdInt, jobIdInt int32) (models.SavedJobsResponse, error) {
	req := &pb.SaveJobRequest{
		UserId: strconv.FormatInt(int64(userIdInt), 10),
		JobId:  strconv.FormatInt(int64(jobIdInt), 10),
	}

	grpcResponse, err := jc.Client.SaveJobs(context.Background(), req)
	if err != nil {
		return models.SavedJobsResponse{}, err
	}

	jobID, err := strconv.ParseInt(grpcResponse.JobId, 10, 64)
	if err != nil {
		return models.SavedJobsResponse{}, err
	}
	userID, err := strconv.ParseInt(grpcResponse.UserId, 10, 64)
	if err != nil {
		return models.SavedJobsResponse{}, err
	}
	savedJobID, err := strconv.ParseInt(grpcResponse.Id, 10, 64)
	if err != nil {
		return models.SavedJobsResponse{}, err
	}

	response := models.SavedJobsResponse{
		ID:          uint(savedJobID),
		JobID:       jobID,
		JobseekerID: userID,
	}
	return response, nil
}

func (jc *jobClient) DeleteSavedJob(jobIdInt, userIdInt int32) error {
	req := &pb.DeleteSavedJobRequest{
		UserId: strconv.FormatInt(int64(userIdInt), 10),
		JobId:  strconv.FormatInt(int64(jobIdInt), 10),
	}

	_, err := jc.Client.DeleteSavedJob(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to delete saved job: %w", err)
	}

	return nil
}

func (jc *jobClient) GetASavedJob(userID int32) ([]models.SavedJobsResponse, error) {
	var savedJobs []models.SavedJobsResponse
	req := &pb.GetSavedJobsRequest{
		UserId: strconv.Itoa(int(userID)),
	}

	grpcResponse, err := jc.Client.GetSavedJobs(context.Background(), req)
	if err != nil {
		return savedJobs, err
	}

	for _, savedJob := range grpcResponse.SavedJobs {
		jobID, err := strconv.ParseInt(savedJob.JobId, 10, 64)
		if err != nil {
			return savedJobs, err
		}
		savedJobID, err := strconv.ParseInt(savedJob.Id, 10, 64)
		if err != nil {
			return savedJobs, err
		}
		jobSeekerId, err := strconv.ParseInt(savedJob.UserId, 10, 64)
		if err != nil {
			return savedJobs, err
		}
		savedJobResponse := models.SavedJobsResponse{
			ID:          uint(savedJobID),
			JobID:       jobID,
			JobseekerID: jobSeekerId,
		}
		savedJobs = append(savedJobs, savedJobResponse)
	}

	fmt.Println("saved jobs", savedJobs)
	return savedJobs, nil
}

// UpdateApplyJob updates the status of a job application
// UpdateApplyJob updates the status of a job application and returns job seeker ID
func (jc *jobClient) UpdateApplyJob(applyJobID uint64, status string) (uint, uint, error) {
    // Construct the request message for updating a job application
    req := &pb.UpdateApplyJobRequest{
        ApplyJobId: uint32(applyJobID),
        Status:     status,
    }

    // Call the gRPC method
    response, err := jc.Client.UpdateApplyJob(context.Background(), req)
    if err != nil {
        log.Printf("gRPC error: %v", err)
        return 0, 0, err
    }
	fmt.Println("jlka",response.JobSeekerId,response.JobId)

    // Debug logging
    log.Printf("gRPC response: %+v", response)

    // Check if the operation was successful
    if !response.Success {
        return 0, 0, fmt.Errorf("failed to update job application status")
    }

    return uint(response.JobSeekerId), uint(response.JobId), nil
}

// GetAcceptedApplicants retrieves the accepted applicants for a given job ID
func (jc *jobClient) GetAcceptedApplicants(jobID int64, status string) ([]models.ApplyJobResponse, error) {
    // Construct the request message for getting accepted applicants
    req := &pb.GetAcceptedApplicantsRequest{
        JobId:  jobID,
        Status: status,
    }

    // Call the gRPC method
    response, err := jc.Client.GetApplicants(context.Background(), req)
    if err != nil {
        log.Printf("gRPC error: %v", err)
        return nil, err
    }

    // Debug logging
    log.Printf("gRPC response: %+v", response)

    if response == nil {
        log.Println("Received nil response from gRPC")
        return nil, fmt.Errorf("received nil response from gRPC")
    }

    // Convert gRPC response to internal model
    var applicants []models.ApplyJobResponse
    for _, app := range response.GetApplicants() {
        applicants = append(applicants, models.ApplyJobResponse{
            ID:          uint(app.Id),
            JobID:       app.JobId,
            JobseekerID: app.JobseekerId,
            CoverLetter: app.CoverLetter,
            ResumeURL:   app.ResumeUrl,
        })
    }

    return applicants, nil
}
