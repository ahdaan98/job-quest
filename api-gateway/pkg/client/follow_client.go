package client

import (
	"context"
	"fmt"

	interfaces "github.com/ahdaan67/JobQuest/pkg/client/interface"
	"github.com/ahdaan67/JobQuest/pkg/config"
	pb "github.com/ahdaan67/JobQuest/pkg/pb/connection"
	"github.com/ahdaan67/JobQuest/pkg/utils/models"
	"google.golang.org/grpc"
)

type followCompanyClient struct {
    client pb.CompanyServiceClient
}

func NewfollowCompanyClient(cfg config.Config) interfaces.FollowCompanyClient {
	grpcConnection, err := grpc.Dial(cfg.JobQuestFollow, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewCompanyServiceClient(grpcConnection)

	return &followCompanyClient{
		client: grpcClient,
	}
}

func (fc *followCompanyClient) FollowCompany(userID int32, companyID int32) (models.FollowCompanyResponse, error) {
    req := &pb.FollowCompanyRequest{
        UserId:    uint32(userID),
        CompanyId: uint32(companyID),
    }
    
    res, err := fc.client.FollowCompany(context.Background(), req)
    if err != nil {
        return models.FollowCompanyResponse{}, err
    }

    return models.FollowCompanyResponse{
        Success: res.Success,
    }, nil
}

func (fc *followCompanyClient) UnfollowCompany(userID int32, companyID int32) (models.UnfollowCompanyResponse, error) {
    req := &pb.UnfollowCompanyRequest{
        UserId:    uint32(userID),
        CompanyId: uint32(companyID),
    }

    res, err := fc.client.UnfollowCompany(context.Background(), req)
    if err != nil {
        return models.UnfollowCompanyResponse{}, err
    }

    return models.UnfollowCompanyResponse{
        Success: res.Success,
    }, nil
}

func (fc *followCompanyClient) IsFollowingCompany(userID int32, companyID int32) (models.IsFollowingCompanyResponse, error) {
    req := &pb.IsFollowingCompanyRequest{
		UserId:    uint32(userID),
        CompanyId: uint32(companyID),
    }

    res, err := fc.client.IsFollowingCompany(context.Background(), req)
    if err != nil {
        return models.IsFollowingCompanyResponse{}, err
    }

    return models.IsFollowingCompanyResponse{
        IsFollowing: res.IsFollowing,
    }, nil
}

func (fc *followCompanyClient) GetFollowedCompanies(userID int32) ([]int32, error) {
    req := &pb.GetFollowedCompaniesRequest{
		UserId:    uint32(userID),
    }

    res, err := fc.client.GetFollowedCompanies(context.Background(), req)
    if err != nil {
        return nil, err
    }

    return res.Id, nil
}

func (fc *followCompanyClient) CheckFollowRequestExists(userID int32, companyID int32) (models.CheckFollowRequestExistsResponse, error) {
    req := &pb.CheckFollowRequestExistsRequest{
        UserId:    uint32(userID),
        CompanyId: uint32(companyID),
    }

    res, err := fc.client.CheckFollowRequestExists(context.Background(), req)
    if err != nil {
        return models.CheckFollowRequestExistsResponse{}, err
    }

    return models.CheckFollowRequestExistsResponse{
        Exists: res.Exists,
    }, nil
}