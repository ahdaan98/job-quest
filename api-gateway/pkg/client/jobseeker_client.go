package client

import (
	"context"
	"fmt"
	"log"

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

func (jc *jobSeekerClient) JobSeekerOTPSignUp(jobSeekerDetails models.JobSeekerSignUp) (models.OtpResponse, error) {
	otpResponse, err := jc.Client.JobSeekerOTPSignUp(context.Background(), &pb.JobSeekerSignupRequest{
		Email:       jobSeekerDetails.Email,
		Password:    jobSeekerDetails.Password,
		FirstName:   jobSeekerDetails.FirstName,
		LastName:    jobSeekerDetails.LastName,
		PhoneNumber: jobSeekerDetails.PhoneNumber,
		DateOfBirth: jobSeekerDetails.DateOfBirth,
		Gender:      jobSeekerDetails.Gender,
	})
	if err != nil {
		return models.OtpResponse{}, err
	}
	return models.OtpResponse{
		Message: otpResponse.Message,
	}, nil
}

// Implement the OTP Verify method
func (jc *jobSeekerClient) JobSeekerVerifyOTP(otpRequest models.JobSeekerVerifyOTPRequest) (models.OtpVerificationResponse, error) {
	// Call gRPC method
	otpVerificationResponse, err := jc.Client.JobSeekerVerifyOTP(context.Background(), &pb.JobSeekerVerifyOTPRequest{
		Email: otpRequest.Email,
		Otp:   otpRequest.Otp,
	})
	if err != nil {
		log.Printf("gRPC error: %v", err)
		return models.OtpVerificationResponse{}, err
	}

	// Debug logging
	log.Printf("gRPC response: %+v", otpVerificationResponse)

	if otpVerificationResponse == nil {
		log.Println("Received nil response from gRPC")
		return models.OtpVerificationResponse{}, fmt.Errorf("received nil response from gRPC")
	}

	if otpVerificationResponse.JobSeekerDetails == nil {
		log.Println("JobSeekerDetails is nil in gRPC response")
		return models.OtpVerificationResponse{
			Status: otpVerificationResponse.Status,
			Token:  otpVerificationResponse.Token,
		}, nil
	}

	// Convert gRPC response to internal model
	jobSeekerDetails := &models.JobSeekerDetailsResponse{
		ID:          uint(otpVerificationResponse.JobSeekerDetails.Id),
		Email:       otpVerificationResponse.JobSeekerDetails.Email,
		FirstName:   otpVerificationResponse.JobSeekerDetails.FirstName,
		LastName:    otpVerificationResponse.JobSeekerDetails.LastName,
		PhoneNumber: otpVerificationResponse.JobSeekerDetails.PhoneNumber,
		DateOfBirth: otpVerificationResponse.JobSeekerDetails.DateOfBirth,
		Gender:      otpVerificationResponse.JobSeekerDetails.Gender,
	}

	return models.OtpVerificationResponse{
		Status:           otpVerificationResponse.Status,
		Token:            otpVerificationResponse.Token,
		JobSeekerDetails: jobSeekerDetails,
	}, nil
}

func (jc *jobSeekerClient) JobSeekerLinkedinSign(jobSeekerDetails models.JobSeekerDetailsResponse) (models.TokenJobSeeker, error) {
	// Construct the request message for LinkedIn sign-in
	req := &pb.JobSeekerLinkedinSignRequest{
		Email:       jobSeekerDetails.Email,
		FirstName:   jobSeekerDetails.FirstName,
		LastName:    jobSeekerDetails.LastName,
		PhoneNumber: jobSeekerDetails.PhoneNumber,
		DateOfBirth: jobSeekerDetails.DateOfBirth,
		Gender:      jobSeekerDetails.Gender,
	}

	// Call the gRPC method
	response, err := jc.Client.JobSeekerLinkedinSign(context.Background(), req)
	if err != nil {
		return models.TokenJobSeeker{}, err
	}

	// Convert the gRPC response to the model type
	return models.TokenJobSeeker{
		JobSeeker: models.JobSeekerDetailsResponse{
			ID:          uint(response.JobSeekerDetails.Id),
			Email:       response.JobSeekerDetails.Email,
			FirstName:   response.JobSeekerDetails.FirstName,
			LastName:    response.JobSeekerDetails.LastName,
			PhoneNumber: response.JobSeekerDetails.PhoneNumber,
			DateOfBirth: response.JobSeekerDetails.DateOfBirth,
			Gender:      response.JobSeekerDetails.Gender,
		},
		Token: response.Token,
	}, nil
}

func (jc *jobSeekerClient) GetEmailByJobSeekerID(id uint) (string, error) {
	// Construct the request message
	req := &pb.GetEmailByJobSeekerIDRequest{
		JobSeekerId: uint64(id),
	}

	// Call the gRPC method
	response, err := jc.Client.GetEmailByJobSeekerID(context.Background(), req)
	if err != nil {
		return "", err
	}

	// Check if the email is empty
	if response.Email == "" {
		return "", fmt.Errorf("email address not found for job seeker ID %d", id)
	}

	return response.Email, nil
}

func (jc *jobSeekerClient) ActivateSubscriptionPlan(jobSeekerID uint, subscriptionPlanID uint) (models.SubscriptionPlanResponse, error) {
	// Construct the request message
	req := &pb.ActivateSubscriptionPlanRequest{
		JobSeekerId: uint32(jobSeekerID),
		PlanId:      uint32(subscriptionPlanID),
	}

	// Call the gRPC method
	response, err := jc.Client.ActivateSubscriptionPlan(context.Background(), req)
	if err != nil {
		return models.SubscriptionPlanResponse{}, err
	}

	// Convert the gRPC response to the model type
	return models.SubscriptionPlanResponse{
		Success: response.Success,
		Message: response.Message,
	}, nil
}
