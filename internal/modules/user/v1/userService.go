package userv1

import (
	"github.com/addixit1/fiber-boilerplate/internal/modules/user"
	"go.mongodb.org/mongo-driver/bson"
)

func ListUsers(filter bson.M) ([]user.User, error) {
	return FindUsers(filter)
}

func CreateUser(user *CreateUserDTO) error {
	_, err := saveUser(user)
	return err
}
