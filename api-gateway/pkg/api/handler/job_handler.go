package handler

import (
	interfaces "github.com/ahdaan67/JobQuest/pkg/client/interface"
	"github.com/ahdaan67/JobQuest/pkg/utils/models"
	"github.com/ahdaan67/JobQuest/pkg/utils/response"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	GRPC_Client interfaces.JobClient
	LogFile     *os.File
}

func NewJobHandler(jobClient interfaces.JobClient) *JobHandler {
	return &JobHandler{
		GRPC_Client: jobClient,
	}
}

func (jh *JobHandler) PostJobOpening(c *gin.Context) {
	employerID, ok := c.Get("id")
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employerIDInt, ok := employerID.(int32)
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	var jobOpening models.JobOpening
	if err := c.ShouldBindJSON(&jobOpening); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	fmt.Println("id", employerIDInt, employerID)

	JobOpening, err := jh.GRPC_Client.PostJobOpening(jobOpening, employerIDInt)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to create job opening", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	response := response.ClientResponse(http.StatusCreated, "Job opening created successfully", JobOpening, nil)
	c.JSON(http.StatusCreated, response)
}

func (jh *JobHandler) GetAllJobs(c *gin.Context) {
	employerID, ok := c.Get("id")
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employerIDInt, ok := employerID.(int32)
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobs, err := jh.GRPC_Client.GetAllJobs(employerIDInt)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) GetAJob(c *gin.Context) {
	idStr := c.Query("id")

	employerID, ok := c.Get("id")
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employerIDInt, ok := employerID.(int32)
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobs, err := jh.GRPC_Client.GetAJob(employerIDInt, int32(jobID))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) DeleteAJob(c *gin.Context) {
	idStr := c.Query("id")

	employerID, ok := c.Get("id")
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employerIDInt, ok := employerID.(int32)
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	err = jh.GRPC_Client.DeleteAJob(employerIDInt, int32(jobID))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to delete job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Job Deleted successfully", nil, nil)
	c.JSON(http.StatusOK, response)
}
func (jh *JobHandler) UpdateAJob(c *gin.Context) {
	idStr := c.Query("id")
	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employerID, ok := c.Get("id")
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employerIDInt, ok := employerID.(int32)
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	var jobOpening models.JobOpening
	if err := c.ShouldBindJSON(&jobOpening); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	UpdateJobOpening, err := jh.GRPC_Client.UpdateAJob(employerIDInt, int32(jobID), jobOpening)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to update job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Job updated successfully", UpdateJobOpening, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) ViewAllJobs(c *gin.Context) {
	keyword := c.Query("Keyword")
	if keyword == "" {
		errs := response.ClientResponse(http.StatusBadRequest, "Keyword parameter is required", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobs, err := jh.GRPC_Client.JobSeekerGetAllJobs(keyword)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	if len(jobs) == 0 {
		errs := response.ClientResponse(http.StatusOK, "No jobs found matching your query", nil, nil)
		c.JSON(http.StatusOK, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) GetJobDetails(c *gin.Context) {
	idStr := c.Query("id")
	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobDetails, err := jh.GRPC_Client.GetJobDetails(int32(jobID))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch job details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Job details retrieved successfully", jobDetails, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) SaveAJob(c *gin.Context) {

	jobIDStr := c.Query("job_id")
	jobIdInt, err := strconv.ParseInt(jobIDStr, 10, 32)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid or missing job ID", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userID, userIDExists := c.Get("id")
	if !userIDExists {
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIdInt, userIDOk := userID.(int32)
	if !userIDOk {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid user ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	Data, err := jh.GRPC_Client.SaveAJob(userIdInt, int32(jobIdInt))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to save job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Job saved successfully", Data, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) DeleteSavedJob(c *gin.Context) {

	jobIDStr := c.Query("job_id")
	jobIdInt, err := strconv.ParseInt(jobIDStr, 10, 32)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userID, userIDExists := c.Get("id")
	if !userIDExists {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid or missing user ID", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIdInt, userIDOk := userID.(int32)
	if !userIDOk {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid user ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	err = jh.GRPC_Client.DeleteSavedJob(int32(jobIdInt), userIdInt)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to delete job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Job deleted successfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) GetASavedJob(c *gin.Context) {

	userID, userIDExists := c.Get("id")
	if !userIDExists {
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIdInt, userIDOk := userID.(int32)
	if !userIDOk {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid user ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}


	job, err := jh.GRPC_Client.GetASavedJob(userIdInt)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to get job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Job fetched successfully", job, nil)
	c.JSON(http.StatusOK, response)
}