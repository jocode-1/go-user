package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/jocode-1/go-user/models"
	"go.mongodb.org/mongo-driver/bson"
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

	_, err := u.usercollection.InsertOne(u.ctx, user)

	return err
}

func (u *UserServiceImpl) GetUser(id *int) (*models.User, error) {

	var user *models.User
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)

	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {

	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {

	filter := bson.D{bson.E{Key: "id", Value: user.Id}}
	update := bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: "id", Value: user.Id}, bson.E{Key: "name", Value: user.Name}, bson.E{Key: "email", Value: user.Email}, bson.E{Key: "address", Value: user.Address}}}

	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return fmt.Sprintf(errors.New("No User Found By%s", filter)) 
	}


	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}
