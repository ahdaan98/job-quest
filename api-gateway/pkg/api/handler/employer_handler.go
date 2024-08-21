package handler

import (
	interfaces "github.com/ahdaan67/JobQuest/pkg/client/interface"
	"github.com/ahdaan67/JobQuest/pkg/utils/models"
	"github.com/ahdaan67/JobQuest/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployerHandler struct {
	GRPC_Client interfaces.EmployerClient
}

func NewEmployerHandler(employerClient interfaces.EmployerClient) *EmployerHandler {
	return &EmployerHandler{
		GRPC_Client: employerClient,
	}
}

func (eh *EmployerHandler) EmployerLogin(c *gin.Context) {

	var employerDetails models.EmployerLogin
	if err := c.ShouldBindJSON(&employerDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employer, err := eh.GRPC_Client.EmployerLogin(employerDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate employer", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Employer authenticated successfully", employer, nil)
	c.JSON(http.StatusOK, success)
}

func (eh *EmployerHandler) EmployerSignUp(c *gin.Context) {

	var employerDetails models.EmployerSignUp
	if err := c.ShouldBindJSON(&employerDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employer, err := eh.GRPC_Client.EmployerSignUp(employerDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot create employer", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Employer created successfully", employer, nil)
	c.JSON(http.StatusOK, success)
}

func (eh *EmployerHandler) GetCompanyDetails(c *gin.Context) {

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

	companyDetails, err := eh.GRPC_Client.GetCompanyDetails(employerIDInt)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch company details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Company details retrieved successfully", companyDetails, nil)
	c.JSON(http.StatusOK, response)
}

func (eh *EmployerHandler) UpdateCompany(c *gin.Context) {

	var employerDetails models.EmployerDetails
	if err := c.ShouldBindJSON(&employerDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
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

	updatedCompany, err := eh.GRPC_Client.UpdateCompany(employerIDInt, employerDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to update company", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Company updated successfully", updatedCompany, nil)
	c.JSON(http.StatusOK, response)
}