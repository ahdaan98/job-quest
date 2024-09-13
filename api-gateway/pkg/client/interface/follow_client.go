package interfaces

import "github.com/ahdaan67/JobQuest/pkg/utils/models"

type FollowCompanyClient interface {
	FollowCompany(userID int32, companyID int32) (models.FollowCompanyResponse, error)
	UnfollowCompany(userID int32, companyID int32) (models.UnfollowCompanyResponse, error)
	IsFollowingCompany(userID int32, companyID int32) (models.IsFollowingCompanyResponse, error)
	GetFollowedCompanies(userID int32) ([]int32, error)
	CheckFollowRequestExists(userID int32, companyID int32) (models.CheckFollowRequestExistsResponse, error)
}