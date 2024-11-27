package repository

import (
	"go_zero-crud/rpc/models"
	"go_zero-crud/rpc/repository/interface_repository"
	"gorm.io/gorm"
)

var _ interface_repository.CompanyDaoInterface = &CompanyRepository{}

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (u *CompanyRepository) FindByEmail(email string) (*models.Company, error) {
	var Company models.Company
	err := u.db.Where("email = ?", email).First(&Company).Error
	if err != nil {
		return nil, err
	}
	return &Company, nil
}

func (u *CompanyRepository) FindByID(id string) (*models.Company, error) {
	var Company models.Company
	err := u.db.Where("id = ?", id).First(&Company).Error
	if err != nil {
		return nil, err
	}
	return &Company, nil
}

func (u *CompanyRepository) Create(Company *models.Company) (*models.Company, error) {
	err := u.db.Create(Company).Error
	if err != nil {
		return nil, err
	}
	return Company, nil
}
