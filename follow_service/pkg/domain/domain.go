package domain

type CompanyFollow struct {
	UserID    uint   `gorm:"primary_key;auto_increment:false"`
	CompanyID uint   `gorm:"primary_key;auto_increment:false"`
	Status    string `gorm:"type:varchar(20);not null"`
}