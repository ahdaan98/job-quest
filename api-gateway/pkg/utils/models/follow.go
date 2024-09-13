package models

type FollowCompanyResponse struct {
    Success bool
    Message string
}

type UnfollowCompanyResponse struct {
    Success bool
    Message string
}

type IsFollowingCompanyResponse struct {
    IsFollowing bool
}

type CheckFollowRequestExistsResponse struct {
    Exists bool
}

type CompanyDetails struct {
    ID          int32
    Name        string
    Description string
}

type FollowCompanyRequest struct {
    CompanyID int32 `json:"company_id"`
}

type UnfollowCompanyRequest struct {
    CompanyID int32 `json:"company_id"`
}