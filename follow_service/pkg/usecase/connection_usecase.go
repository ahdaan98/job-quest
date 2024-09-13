package usecase

import (
	"errors"
	"fmt"
	interfaces "follow/pkg/repository/interface"
	services "follow/pkg/usecase/interface"
)

type companyUseCase struct {
	companyRepository interfaces.ConnectionRepository
}

func NewCompanyUseCase(repository interfaces.ConnectionRepository) services.CompanyUseCase {
	return &companyUseCase{
		companyRepository: repository,
	}
}

func (cu *companyUseCase) FollowCompany(userID, companyID uint) error {
	if userID <= 0 || companyID <= 0 {
		return errors.New("invalid input data")
	}

	err := cu.companyRepository.FollowCompany(userID, companyID)
	if err != nil {
		return fmt.Errorf("failed to follow company: %v", err)
	}

	return nil
}

func (cu *companyUseCase) UnfollowCompany(userID, companyID uint) error {
	if userID <= 0 || companyID <= 0 {
		return errors.New("invalid input data")
	}

	err := cu.companyRepository.UnfollowCompany(userID, companyID)
	if err != nil {
		return fmt.Errorf("failed to unfollow company: %v", err)
	}

	return nil
}

func (cu *companyUseCase) IsFollowingCompany(userID, companyID uint) (bool, error) {
	if userID <= 0 || companyID <= 0 {
		return false, errors.New("invalid input data")
	}

	isFollowing, err := cu.companyRepository.IsFollowingCompany(userID, companyID)
	if err != nil {
		return false, fmt.Errorf("failed to check if following company: %v", err)
	}

	return isFollowing, nil
}

func (cu *companyUseCase) GetFollowedCompanies(userID uint) ([]int, error) {
	if userID <= 0 {
		return nil, errors.New("invalid input data")
	}

	companies, err := cu.companyRepository.GetFollowedCompanies(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get followed companies: %v", err)
	}

	return companies, nil
}

func (cu *companyUseCase) CheckFollowRequestExists(userID, companyID uint) (bool, error) {
	if userID <= 0 || companyID <= 0 {
		return false, errors.New("invalid input data")
	}

	exists, err := cu.companyRepository.CheckFollowRequestExists(userID, companyID)
	if err != nil {
		return false, fmt.Errorf("failed to check follow request existence: %v", err)
	}

	return exists, nil
}
