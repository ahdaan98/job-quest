package repository

import (
	"errors"

	"follow/pkg/domain"
	interfaces "follow/pkg/repository/interface"
	"gorm.io/gorm"
)

type ConnectionRepository struct {
	DB          *gorm.DB
}

func NewConnectionRepository(DB *gorm.DB) interfaces.ConnectionRepository {
	return &ConnectionRepository{
		DB:          DB,
	}
}

func (ur *ConnectionRepository) FollowCompany(userID, companyID uint) error {
	connection := domain.CompanyFollow{
		UserID:    userID,
		CompanyID: companyID,
		Status:    "following",
	}
	result := ur.DB.Where("user_id = ? AND company_id = ?", userID, companyID).First(&connection)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	if result.RowsAffected > 0 {
		connection.Status = "following"
		result = ur.DB.Save(&connection)
	} else {
		result = ur.DB.Create(&connection)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *ConnectionRepository) UnfollowCompany(userID, companyID uint) error {
	result := ur.DB.Where("user_id = ? AND company_id = ?", userID, companyID).Delete(&domain.CompanyFollow{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *ConnectionRepository) IsFollowingCompany(userID, companyID uint) (bool, error) {
	var count int64
	err := ur.DB.Table("company_follows").
		Where("user_id = ? AND company_id = ? AND status = 'following'", userID, companyID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ur *ConnectionRepository) GetFollowedCompanies(userID uint) ([]int, error) {
	query := `
	SELECT company_id FROM company_follows WHERE user_id=$1; 
	`
	var ids []int
	if err := ur.DB.Raw(query,userID).Scan(&ids).Error; err!=nil {
		return nil, err
	}

	return ids, nil
}

func (ur *ConnectionRepository) CheckFollowRequestExists(userID, companyID uint) (bool, error) {
	var count int64
	err := ur.DB.Table("company_follows").
		Where("user_id = ? AND company_id = ? AND status = 'pending'", userID, companyID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
