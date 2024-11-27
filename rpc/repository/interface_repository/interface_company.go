package interface_repository

import "go_zero-crud/rpc/models"

type CompanyDaoInterface interface {
	FindByEmail(email string) (*models.Company, error)
	FindByID(id string) (*models.Company, error)
	Create(user *models.Company) (*models.Company, error)
}
