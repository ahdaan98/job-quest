package handler

import (
	"errors"
	"net/http"
	"strconv"

	interfaces "github.com/ahdaan67/JobQuest/pkg/client/interface"
	"github.com/ahdaan67/JobQuest/pkg/utils/models"
	"github.com/ahdaan67/JobQuest/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

type FollowCompanyHandler struct {
	GRPC_Client interfaces.FollowCompanyClient
	e           interfaces.EmployerClient
}

func NewFollowCompanyHandler(followCompanyClient interfaces.FollowCompanyClient, e interfaces.EmployerClient) *FollowCompanyHandler {
	return &FollowCompanyHandler{
		GRPC_Client: followCompanyClient,
		e:           e,
	}
}

func (fh *FollowCompanyHandler) FollowCompany(c *gin.Context) {
	userID, userIDExists := c.Get("id")
	if !userIDExists {
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Assuming userID is a string or something similar that needs to be converted to an int32
	userIDStr, ok := userID.(uint)
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid User ID format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIDInt32 := int32(userIDStr)

	var followRequest models.FollowCompanyRequest
	if err := c.ShouldBindJSON(&followRequest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid request format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	res, err := fh.e.GetCompanyDetails(followRequest.CompanyID)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to follow company", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	if res.Company_name == "" {
		errs := response.ClientResponse(http.StatusInternalServerError, "Company does not exist", nil, errors.New("company does not exist"))
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	result, err := fh.GRPC_Client.FollowCompany(userIDInt32, followRequest.CompanyID)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to follow company", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	result.Message = res.Company_name

	success := response.ClientResponse(http.StatusOK, "Successfully followed the company", result, nil)
	c.JSON(http.StatusOK, success)
}

func (fh *FollowCompanyHandler) UnfollowCompany(c *gin.Context) {
	userID, userIDExists := c.Get("id")
	if !userIDExists {
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Assuming userID is a string or something similar that needs to be converted to an int32
	userIDStr, ok := userID.(uint)
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid User ID format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIDInt64:= userIDStr

	userIDInt32 := int32(userIDInt64)

	var unfollowRequest models.UnfollowCompanyRequest
	if err := c.ShouldBindJSON(&unfollowRequest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid request format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	result, err := fh.GRPC_Client.UnfollowCompany(userIDInt32, unfollowRequest.CompanyID)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to unfollow company", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Successfully unfollowed the company", result, nil)
	c.JSON(http.StatusOK, success)
}

func (fh *FollowCompanyHandler) IsFollowingCompany(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid user ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	companyID, err := strconv.Atoi(c.Param("companyID"))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid company ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	result, err := fh.GRPC_Client.IsFollowingCompany(int32(userID), int32(companyID))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to check follow status", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Follow status retrieved", result, nil)
	c.JSON(http.StatusOK, success)
}

func (fh *FollowCompanyHandler) GetFollowedCompanies(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid user ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	companies, err := fh.GRPC_Client.GetFollowedCompanies(int32(userID))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to retrieve followed companies", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	var cf []models.FEmployer
	var cr models.FEmployer
	for _,v  := range companies{
		k, err := fh.e.GetCompanyDetails(v)
		if err != nil {
			errs := response.ClientResponse(http.StatusInternalServerError, "Failed to get followed companies details", nil, err.Error())
			c.JSON(http.StatusInternalServerError, errs)
			return
		}

		cr.ID = k.ID
		cr.Company_name = k.Company_name

		cf = append(cf, cr)
	}

	success := response.ClientResponse(http.StatusOK, "Followed companies retrieved successfully", cf, nil)
	c.JSON(http.StatusOK, success)
}

func (fh *FollowCompanyHandler) CheckFollowRequestExists(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid user ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	companyID, err := strconv.Atoi(c.Param("companyID"))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid company ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	result, err := fh.GRPC_Client.CheckFollowRequestExists(int32(userID), int32(companyID))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to check follow request existence", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Follow request existence retrieved", result, nil)
	c.JSON(http.StatusOK, success)
}
