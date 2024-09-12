// job_seeker_service.go
package service

import (
	"Auth/pkg/config"
	"Auth/pkg/helper"
	pb "Auth/pkg/pb/auth"
	interfaces "Auth/pkg/usecase/interface"
	"Auth/pkg/utils/models"
	"context"
	"log"
	"strconv"
	"sync"
)

type JobSeekerServer struct {
	jobSeekerUseCase interfaces.JobSeekerUseCase
	signupDetailsMap map[string]models.JobSeekerSignUp
	otpStore         map[string]string // Added for OTP storage
	mu               sync.Mutex
	pb.UnimplementedJobSeekerServer
}

func NewJobSeekerServer(useCase interfaces.JobSeekerUseCase) pb.JobSeekerServer {
	return &JobSeekerServer{
		jobSeekerUseCase: useCase,
		signupDetailsMap: make(map[string]models.JobSeekerSignUp),
		otpStore:         make(map[string]string), // Initialize OTP store
	}
}

func (js *JobSeekerServer) JobSeekerSignup(ctx context.Context, req *pb.JobSeekerSignupRequest) (*pb.JobSeekerSignupResponse, error) {
	jobSeekerSignup := models.JobSeekerSignUp{
		Email:       req.Email,
		Password:    req.Password,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
	}

	log.Printf("Received signup request: %+v", jobSeekerSignup)

	res, err := js.jobSeekerUseCase.JobSeekerSignUp(jobSeekerSignup)
	if err != nil {
		return &pb.JobSeekerSignupResponse{}, err
	}

	jobSeekerDetails := &pb.JobSeekerDetails{
		Id:          uint64(res.JobSeeker.ID),
		Email:       res.JobSeeker.Email,
		FirstName:   res.JobSeeker.FirstName,
		LastName:    res.JobSeeker.LastName,
		PhoneNumber: res.JobSeeker.PhoneNumber,
		DateOfBirth: res.JobSeeker.DateOfBirth,
		Gender:      res.JobSeeker.Gender,
	}

	return &pb.JobSeekerSignupResponse{
		Status:           201,
		JobSeekerDetails: jobSeekerDetails,
		Token:            res.Token,
	}, nil
}

func (js *JobSeekerServer) JobSeekerLogin(ctx context.Context, req *pb.JobSeekerLoginRequest) (*pb.JobSeekerLoginResponse, error) {
	jobSeekerLogin := models.JobSeekerLogin{
		Email:    req.Email,
		Password: req.Password,
	}

	jobSeeker, err := js.jobSeekerUseCase.JobSeekerLogin(jobSeekerLogin)
	if err != nil {
		return &pb.JobSeekerLoginResponse{}, err
	}

	jobSeekerDetails := &pb.JobSeekerDetails{
		Id:          uint64(jobSeeker.JobSeeker.ID),
		Email:       jobSeeker.JobSeeker.Email,
		FirstName:   jobSeeker.JobSeeker.FirstName,
		LastName:    jobSeeker.JobSeeker.LastName,
		PhoneNumber: jobSeeker.JobSeeker.PhoneNumber,
		DateOfBirth: jobSeeker.JobSeeker.DateOfBirth,
		Gender:      jobSeeker.JobSeeker.Gender,
	}

	return &pb.JobSeekerLoginResponse{
		Status:           200,
		JobSeekerDetails: jobSeekerDetails,
		Token:            jobSeeker.Token,
	}, nil
}

func (js *JobSeekerServer) JobSeekerOTPSignUp(ctx context.Context, req *pb.JobSeekerSignupRequest) (*pb.JobSeekerOTPSignUpResponse, error) {
	// Validate user details (e.g., check if the user already exists)
	jobSeekerSignup := models.JobSeekerSignUp{
		Email:       req.Email,
		Password:    req.Password,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
	}

	log.Printf("Received signup request: %+v", jobSeekerSignup)

	// Load configuration for email
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %v", err)
		return &pb.JobSeekerOTPSignUpResponse{
			Message: "Internal server error: " + err.Error(),
		}, nil
	}

	// Generate and send OTP
	otpStr, err := helper.SendOTP(req.Email, cfg)
	if err != nil {
		log.Printf("Failed to send OTP: %v", err)
		return &pb.JobSeekerOTPSignUpResponse{
			Message: "Failed to send OTP: " + err.Error(),
		}, nil
	}

	// Store OTP and signup details
	js.mu.Lock()
	defer js.mu.Unlock()
	js.otpStore[req.Email] = otpStr
	js.signupDetailsMap[req.Email] = jobSeekerSignup

	log.Printf("Sent OTP: %s to email: %s", otpStr, req.Email)

	// Successfully sent OTP, respond with a message
	return &pb.JobSeekerOTPSignUpResponse{
		Message: "OTP sent successfully to " + req.Email + ". Please verify your OTP.",
	}, nil
}

func (js *JobSeekerServer) JobSeekerVerifyOTP(ctx context.Context, req *pb.JobSeekerVerifyOTPRequest) (*pb.JobSeekerSignupResponse, error) {
	// Convert OTP from int32 to string
	otpStr := strconv.Itoa(int(req.Otp))
	log.Printf("Received OTP: %s for email: %s", otpStr, req.Email)

	// Verify the OTP
	js.mu.Lock()
	storedOtp, exists := js.otpStore[req.Email]
	js.mu.Unlock()

	if !exists || storedOtp != otpStr {
		log.Printf("Invalid OTP for email: %s", req.Email)
		return &pb.JobSeekerSignupResponse{
			Status: 400,
			// No details or token in case of invalid OTP
		}, nil
	}
	log.Printf("OTP verified successfully for email: %s", req.Email)

	// Retrieve the stored signup details
	js.mu.Lock()
	defer js.mu.Unlock()
	jobSeekerSignup, exists := js.signupDetailsMap[req.Email]
	if !exists {
		log.Printf("Signup details not found for email: %s", req.Email)
		return &pb.JobSeekerSignupResponse{
			Status: 400,
			// No details or token if the user is not found
		}, nil
	}
	log.Printf("Signup details found for email: %s", req.Email)

	// Complete the signup process
	res, err := js.jobSeekerUseCase.JobSeekerSignUp(jobSeekerSignup)
	if err != nil {
		log.Printf("Error during signup process for email: %s, error: %v", req.Email, err)
		return &pb.JobSeekerSignupResponse{
			Status: 500,
			// No details or token in case of error
		}, nil
	}
	log.Printf("Signup completed successfully for email: %s", req.Email)

	// Prepare the response with job seeker details and token
	jobSeekerDetails := &pb.JobSeekerDetails{
		Id:          uint64(res.JobSeeker.ID),
		Email:       res.JobSeeker.Email,
		FirstName:   res.JobSeeker.FirstName,
		LastName:    res.JobSeeker.LastName,
		PhoneNumber: res.JobSeeker.PhoneNumber,
		DateOfBirth: res.JobSeeker.DateOfBirth,
		Gender:      res.JobSeeker.Gender,
	}

	log.Printf("Returning response for email: %s with status: 201", req.Email)
	return &pb.JobSeekerSignupResponse{
		Status:           200,
		JobSeekerDetails: jobSeekerDetails,
		Token:            res.Token,
	}, nil
}

func (js *JobSeekerServer) JobSeekerLinkedinSign(ctx context.Context, req *pb.JobSeekerLinkedinSignRequest) (*pb.JobSeekerSignupResponse, error) {
	// Map request to domain model
	jobSeekerDetails := models.JobSeekerDetailsResponse{
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
	}

	// Log received request
	log.Printf("Received LinkedIn sign-in request: %+v", jobSeekerDetails)

	// Call the use case method
	tokenJobSeeker, err := js.jobSeekerUseCase.JobSeekerLinkedinSign(jobSeekerDetails)
	if err != nil {
		log.Printf("Error during LinkedIn sign-in: %v", err)
		return &pb.JobSeekerSignupResponse{
			Status: 500,
		}, nil
	}

	// Prepare response with job seeker details and token
	jobSeekerResponse := &pb.JobSeekerDetails{
		Id:          uint64(tokenJobSeeker.JobSeeker.ID),
		Email:       tokenJobSeeker.JobSeeker.Email,
		FirstName:   tokenJobSeeker.JobSeeker.FirstName,
		LastName:    tokenJobSeeker.JobSeeker.LastName,
		PhoneNumber: tokenJobSeeker.JobSeeker.PhoneNumber,
		DateOfBirth: tokenJobSeeker.JobSeeker.DateOfBirth,
		Gender:      tokenJobSeeker.JobSeeker.Gender,
	}

	// Log response preparation
	log.Printf("Returning response for LinkedIn sign-in: %+v", jobSeekerResponse)

	return &pb.JobSeekerSignupResponse{
		Status:           200,
		JobSeekerDetails: jobSeekerResponse,
		Token:            tokenJobSeeker.Token,
	}, nil
}

func (js *JobSeekerServer) GetEmailByJobSeekerID(ctx context.Context, req *pb.GetEmailByJobSeekerIDRequest) (*pb.GetEmailByJobSeekerIDResponse, error) {
	// Extract job seeker ID from the request
	jobSeekerID := req.GetJobSeekerId()

	// Log the received request for debugging
	log.Printf("Received request to get email by job seeker ID: %d", jobSeekerID)

	// Call the use case to retrieve the email by job seeker ID
	email, _ := js.jobSeekerUseCase.GetEmailByJobSeekerID(uint(jobSeekerID))

	// Check if the email was found or not
	if email == "" {
		log.Printf("No email found for job seeker ID: %d", jobSeekerID)
		return &pb.GetEmailByJobSeekerIDResponse{
			Email: "",
		}, nil
	}

	// Log success and prepare response
	log.Printf("Email found for job seeker ID %d: %s", jobSeekerID, email)

	return &pb.GetEmailByJobSeekerIDResponse{
		Email: email,
	}, nil
}

func (js *JobSeekerServer) ActivateSubscriptionPlan(ctx context.Context, req *pb.ActivateSubscriptionPlanRequest) (*pb.ActivateSubscriptionPlanResponse, error) {
	// Extract job seeker ID and subscription plan ID from the request
	jobSeekerID := req.GetJobSeekerId()
	subscriptionPlanID := req.GetPlanId()

	// Log the received request for debugging
	log.Printf("Received request to activate subscription plan: JobSeekerID=%d, SubscriptionPlanID=%d", jobSeekerID, subscriptionPlanID)

	// Call the use case to activate the subscription plan
	message, err := js.jobSeekerUseCase.ActivateSubscriptionPlan(uint(jobSeekerID), uint(subscriptionPlanID))
	if err != nil {
		// Log the error and prepare the response with failure status
		log.Printf("Error activating subscription plan: %v", err)
		return &pb.ActivateSubscriptionPlanResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	// Log success and prepare the response
	log.Printf("Successfully activated subscription plan for JobSeekerID=%d", jobSeekerID)

	return &pb.ActivateSubscriptionPlanResponse{
		Success: true,
		Message: message,
	}, nil
}
