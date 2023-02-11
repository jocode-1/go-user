package services

import "github.com/jocode-1/go-user/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
