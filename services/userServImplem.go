package services

import (
	"context"

	"github.com/jocode-1/go-user/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {

	return nil
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {

	return nil, nil
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {

	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {

	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}
