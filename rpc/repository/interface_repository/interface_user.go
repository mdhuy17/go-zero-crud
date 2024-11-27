package interface_repository

import "go_zero-crud/rpc/models"

type UserDaoInterface interface {
	FindByEmail(email string) (*models.User, error)
	FindByID(id string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
}
