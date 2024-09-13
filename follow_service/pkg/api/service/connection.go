package service

import (
	"context"
	pb "follow/pkg/pb/connection"
	interfaces "follow/pkg/usecase/interface"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CompanyServer struct {
	companyUseCase interfaces.CompanyUseCase
	pb.UnimplementedCompanyServiceServer
}

func NewCompanyServer(useCase interfaces.CompanyUseCase) pb.CompanyServiceServer {
	return &CompanyServer{
		companyUseCase: useCase,
	}
}

func (cs *CompanyServer) FollowCompany(ctx context.Context, req *pb.FollowCompanyRequest) (*pb.FollowCompanyResponse, error) {
	userID := req.UserId
	companyID := req.CompanyId

	err := cs.companyUseCase.FollowCompany(uint(userID), uint(companyID))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to follow company: %v", err)
	}

	return &pb.FollowCompanyResponse{Success: true}, nil
}

func (cs *CompanyServer) UnfollowCompany(ctx context.Context, req *pb.UnfollowCompanyRequest) (*pb.UnfollowCompanyResponse, error) {
	userID := req.UserId
	companyID := req.CompanyId

	err := cs.companyUseCase.UnfollowCompany(uint(userID), uint(companyID))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unfollow company: %v", err)
	}

	return &pb.UnfollowCompanyResponse{Success: true}, nil
}

func (cs *CompanyServer) IsFollowingCompany(ctx context.Context, req *pb.IsFollowingCompanyRequest) (*pb.IsFollowingCompanyResponse, error) {
	userID := req.UserId
	companyID := req.CompanyId

	isFollowing, err := cs.companyUseCase.IsFollowingCompany(uint(userID), uint(companyID))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check if following company: %v", err)
	}

	return &pb.IsFollowingCompanyResponse{IsFollowing: isFollowing}, nil
}

func (cs *CompanyServer) GetFollowedCompanies(ctx context.Context, req *pb.GetFollowedCompaniesRequest) (*pb.GetFollowedCompaniesResponse, error) {
	userID := req.UserId

	followedCompanies, err := cs.companyUseCase.GetFollowedCompanies(uint(userID))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get followed companies: %v", err)
	}

	followedCompaniesInt32 := make([]int32, len(followedCompanies))
    for i, companyID := range followedCompanies {
        followedCompaniesInt32[i] = int32(companyID)
    }

	return &pb.GetFollowedCompaniesResponse{Id: followedCompaniesInt32}, nil
}

func (cs *CompanyServer) CheckFollowRequestExists(ctx context.Context, req *pb.CheckFollowRequestExistsRequest) (*pb.CheckFollowRequestExistsResponse, error) {
	userID := req.UserId
	companyID := req.CompanyId

	exists, err := cs.companyUseCase.CheckFollowRequestExists(uint(userID), uint(companyID))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check if follow request exists: %v", err)
	}

	return &pb.CheckFollowRequestExistsResponse{Exists: exists}, nil
}
