package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	interfaces "github.com/ahdaan67/JobQuest/pkg/client/interface"
	"github.com/ahdaan67/JobQuest/pkg/config"
	"github.com/ahdaan67/JobQuest/pkg/utils/models"
	"github.com/ahdaan67/JobQuest/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

type JobSeekerHandler struct {
	GRPC_Client interfaces.JobSeekerClient
}

func NewJobSeekerHandler(jobSeekerClient interfaces.JobSeekerClient) *JobSeekerHandler {
	return &JobSeekerHandler{
		GRPC_Client: jobSeekerClient,
	}
}

func (jh *JobSeekerHandler) JobSeekerLogin(c *gin.Context) {

	var jobSeekerDetails models.JobSeekerLogin
	if err := c.ShouldBindJSON(&jobSeekerDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobSeeker, err := jh.GRPC_Client.JobSeekerLogin(jobSeekerDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate job seeker", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Job seeker authenticated successfully", jobSeeker, nil)
	c.JSON(http.StatusOK, success)
}

func (jh *JobSeekerHandler) JobSeekerSignUp(c *gin.Context) {

	var jobSeekerDetails models.JobSeekerSignUp
	if err := c.ShouldBindJSON(&jobSeekerDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobSeeker, err := jh.GRPC_Client.JobSeekerSignUp(jobSeekerDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot create job seeker", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Job seeker created successfully", jobSeeker, nil)
	c.JSON(http.StatusOK, success)
}

func (jh *JobSeekerHandler) JobSeekerOTPSignUp(c *gin.Context) {
	var jobSeekerDetails models.JobSeekerSignUp
	if err := c.ShouldBindJSON(&jobSeekerDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Call gRPC method to send OTP
	otpResponse, err := jh.GRPC_Client.JobSeekerOTPSignUp(jobSeekerDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to send OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	// Respond with the OTP message and job seeker details if needed
	success := response.ClientResponse(http.StatusOK, otpResponse.Message, nil, nil)
	c.JSON(http.StatusOK, success)
}

// JobSeekerVerifyOTP handles OTP verification requests.
func (jh *JobSeekerHandler) JobSeekerVerifyOTP(c *gin.Context) {
	var otpRequest models.JobSeekerVerifyOTPRequest
	if err := c.ShouldBindJSON(&otpRequest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Call gRPC method to verify OTP
	verifyResponse, err := jh.GRPC_Client.JobSeekerVerifyOTP(otpRequest)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "OTP verification failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	if verifyResponse.Status == 200 {
		success := response.ClientResponse(http.StatusOK, "OTP verified successfully", verifyResponse.JobSeekerDetails, verifyResponse.Token)
		c.JSON(http.StatusOK, success)
	} else {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid OTP", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
	}
}

func (jh *JobSeekerHandler) LinkedinSignIn(c *gin.Context) {
	url := config.AppConfig.LinkedInLoginConfig.AuthCodeURL("randomstate")
	c.Redirect(http.StatusSeeOther, url)
}

func fetchWithBearerToken(url, token string) (*http.Response, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    // Set the Authorization header with the Bearer token
    req.Header.Set("Authorization", "Bearer " + token)
    client := &http.Client{}
    return client.Do(req)
}

func (jh *JobSeekerHandler) CompleteLinkedInSignIn(c *gin.Context) {
    // Get the access token from the query parameters
    accessToken := c.Query("access_token")
    if accessToken == "" {
        errs := response.ClientResponse(http.StatusBadRequest, "Missing access token", nil, nil)
        c.JSON(http.StatusBadRequest, errs)
        return
    }

    // Define the LinkedIn API URL
    url := "https://api.linkedin.com/v2/userinfo"

    // Fetch user profile data from LinkedIn API
    resp, err := fetchWithBearerToken(url, accessToken)
    if err != nil {
        errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch user data", nil, err.Error())
        c.JSON(http.StatusInternalServerError, errs)
        return
    }
    defer resp.Body.Close()

    // Check for HTTP status codes
    if resp.StatusCode != http.StatusOK {
        errs := response.ClientResponse(http.StatusBadRequest, "Failed to fetch user data", nil, resp.Status)
        c.JSON(http.StatusBadRequest, errs)
        return
    }

    // Read and parse user data
    userData, err := io.ReadAll(resp.Body)
    if err != nil {
        errs := response.ClientResponse(http.StatusInternalServerError, "Failed to read response body", nil, err.Error())
        c.JSON(http.StatusInternalServerError, errs)
        return
    }

    var userInfo map[string]interface{}
    if err := json.Unmarshal(userData, &userInfo); err != nil {
        errs := response.ClientResponse(http.StatusInternalServerError, "Failed to parse user data", nil, err.Error())
        c.JSON(http.StatusInternalServerError, errs)
        return
    }

    // Extract user information
    email, _ := userInfo["email"].(string)
    givenName, _ := userInfo["firstName"].(string)
    familyName, _ := userInfo["lastName"].(string)

    // Create the request object for gRPC
    grpcRequest := models.JobSeekerDetailsResponse{
        Email:       email,
        FirstName:   givenName,
        LastName:    familyName,
    }

    // Call gRPC method to complete LinkedIn sign-in
    grpcResponse, err := jh.GRPC_Client.JobSeekerLinkedinSign(grpcRequest)
    if err != nil {
        errs := response.ClientResponse(http.StatusInternalServerError, "LinkedIn sign-in failed", nil, err.Error())
        c.JSON(http.StatusInternalServerError, errs)
        return
    }

    // Create the response object for success
    success := response.ClientResponse(http.StatusOK, "LinkedIn sign-in completed successfully", grpcResponse, nil)
    c.JSON(http.StatusOK, success)
}

func (jh *JobSeekerHandler) ActivateSubscriptionPlan(c *gin.Context) {
    // Extract job seeker ID and subscription plan ID from the query parameters
    jobSeekerIDStr, userIDExists := c.Get("id")
	if !userIDExists {
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

    subscriptionPlanIDStr := c.Query("plan_id")
    
    if jobSeekerIDStr == "" || subscriptionPlanIDStr == "" {
        errs := response.ClientResponse(http.StatusBadRequest, "Missing job seeker ID or subscription plan ID", nil, nil)
        c.JSON(http.StatusBadRequest, errs)
        return
    }

    jobSeekerID, err := strconv.ParseUint(fmt.Sprintf("%v", jobSeekerIDStr), 10, 32)
    if err != nil {
        errs := response.ClientResponse(http.StatusBadRequest, "Invalid job seeker ID", nil, err.Error())
        c.JSON(http.StatusBadRequest, errs)
        return
    }

    subscriptionPlanID, err := strconv.ParseUint(subscriptionPlanIDStr, 10, 32)
    if err != nil {
        errs := response.ClientResponse(http.StatusBadRequest, "Invalid subscription plan ID", nil, err.Error())
        c.JSON(http.StatusBadRequest, errs)
        return
    }

    // Call gRPC method to activate the subscription plan
    grpcResponse, err := jh.GRPC_Client.ActivateSubscriptionPlan(uint(jobSeekerID), uint(subscriptionPlanID))
    if err != nil {
        errs := response.ClientResponse(http.StatusInternalServerError, "Failed to activate subscription plan", nil, err.Error())
        c.JSON(http.StatusInternalServerError, errs)
        return
    }

    // Check if the activation was successful
    if !grpcResponse.Success {
        errs := response.ClientResponse(http.StatusBadRequest, "Failed to activate subscription plan", nil, grpcResponse.Message)
        c.JSON(http.StatusBadRequest, errs)
        return
    }

    // Create the response object for success
    success := response.ClientResponse(http.StatusOK, "Subscription plan activated successfully", nil, nil)
    c.JSON(http.StatusOK, success)
}