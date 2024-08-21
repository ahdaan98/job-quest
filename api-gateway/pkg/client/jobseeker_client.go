package client

import (
	"context"
	"fmt"

	interfaces "github.com/ahdaan67/JobQuest/pkg/client/interface"
	"github.com/ahdaan67/JobQuest/pkg/config"
	pb "github.com/ahdaan67/JobQuest/pkg/pb/auth"
	"github.com/ahdaan67/JobQuest/pkg/utils/models"

	"google.golang.org/grpc"
)

type jobSeekerClient struct {
	Client pb.JobSeekerClient
}

func NewJobSeekerClient(cfg config.Config) interfaces.JobSeekerClient {
	grpcConnection, err := grpc.Dial(cfg.JobQuestAuth, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewJobSeekerClient(grpcConnection)

	return &jobSeekerClient{
		Client: grpcClient,
	}
}

func (jc *jobSeekerClient) JobSeekerSignUp(jobSeekerDetails models.JobSeekerSignUp) (models.TokenJobSeeker, error) {
	jobSeeker, err := jc.Client.JobSeekerSignup(context.Background(), &pb.JobSeekerSignupRequest{
		Email:       jobSeekerDetails.Email,
		Password:    jobSeekerDetails.Password,
		FirstName:   jobSeekerDetails.FirstName,
		LastName:    jobSeekerDetails.LastName,
		PhoneNumber: jobSeekerDetails.PhoneNumber,
		DateOfBirth: jobSeekerDetails.DateOfBirth,
		Gender:      jobSeekerDetails.Gender,
	})
	if err != nil {
		return models.TokenJobSeeker{}, err
	}
	return models.TokenJobSeeker{
		JobSeeker: models.JobSeekerDetailsResponse{
			ID:          uint(jobSeeker.JobSeekerDetails.Id),
			Email:       jobSeeker.JobSeekerDetails.Email,
			FirstName:   jobSeeker.JobSeekerDetails.FirstName,
			LastName:    jobSeeker.JobSeekerDetails.LastName,
			PhoneNumber: jobSeeker.JobSeekerDetails.PhoneNumber,
			DateOfBirth: jobSeeker.JobSeekerDetails.DateOfBirth,
			Gender:      jobSeeker.JobSeekerDetails.Gender,
		},
		Token: jobSeeker.Token,
	}, nil
}

func (jc *jobSeekerClient) JobSeekerLogin(jobSeekerDetails models.JobSeekerLogin) (models.TokenJobSeeker, error) {
	jobSeeker, err := jc.Client.JobSeekerLogin(context.Background(), &pb.JobSeekerLoginRequest{
		Email:    jobSeekerDetails.Email,
		Password: jobSeekerDetails.Password,
	})

	if err != nil {
		return models.TokenJobSeeker{}, err
	}
	return models.TokenJobSeeker{
		JobSeeker: models.JobSeekerDetailsResponse{
			ID:          uint(jobSeeker.JobSeekerDetails.Id),
			Email:       jobSeeker.JobSeekerDetails.Email,
			FirstName:   jobSeeker.JobSeekerDetails.FirstName,
			LastName:    jobSeeker.JobSeekerDetails.LastName,
			PhoneNumber: jobSeeker.JobSeekerDetails.PhoneNumber,
			DateOfBirth: jobSeeker.JobSeekerDetails.DateOfBirth,
			Gender:      jobSeeker.JobSeekerDetails.Gender,
		},
		Token: jobSeeker.Token,
	}, nil
}