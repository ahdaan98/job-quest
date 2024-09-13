package interfaces

type ConnectionRepository interface {
	FollowCompany(userID, companyID uint) error
	UnfollowCompany(userID, companyID uint) error
	IsFollowingCompany(userID, companyID uint) (bool, error)
	GetFollowedCompanies(userID uint) ([]int, error)
	CheckFollowRequestExists(userID, companyID uint) (bool, error)
}