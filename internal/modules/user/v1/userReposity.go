package userv1

import (
	"context"

	"github.com/addixit1/fiber-boilerplate/internal/modules/user"
	"github.com/addixit1/fiber-boilerplate/internal/querybuilder"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var repo = querybuilder.NewBaseRepository()

// FindUsers retrieves users based on filter using base repository
func FindUsers(filter bson.M) ([]user.User, error) {
	ctx := context.Background()
	users := []user.User{}

	// Using base repository's Find method
	err := repo.Find(ctx, &user.User{}, &users, filter, nil)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FindUserById retrieves a user by ID
func FindUserById(id string) (*user.User, error) {
	ctx := context.Background()

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	foundUser := &user.User{}
	err = repo.FindById(ctx, foundUser, objectID)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

// FindUserByEmail retrieves a user by email
func FindUserByEmail(email string) (*user.User, error) {
	ctx := context.Background()
	foundUser := &user.User{}

	filter := bson.M{"email": email}
	err := repo.FindOne(ctx, foundUser, filter, nil)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

// saveUser creates a new user using base repository
func saveUser(userData *CreateUserDTO) (*user.User, error) {
	ctx := context.Background()

	// Create new User instance
	newUser := &user.User{
		Name:  userData.Name,
		Email: userData.Email,
	}

	// Save using base repository
	err := repo.Save(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// UpdateUser updates an existing user
func UpdateUser(id string, updateData bson.M) error {
	ctx := context.Background()

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updateData}

	_, err = repo.UpdateOne(ctx, &user.User{}, filter, update)
	return err
}

// DeleteUser deletes a user by ID
func DeleteUser(id string) error {
	ctx := context.Background()

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	_, err = repo.DeleteOne(ctx, &user.User{}, filter)
	return err
}

// CountUsers counts users matching the filter
func CountUsers(filter bson.M) (int64, error) {
	ctx := context.Background()
	return repo.CountDocuments(ctx, &user.User{}, filter)
}

// FindUsersWithPagination retrieves users with pagination
func FindUsersWithPagination(filter bson.M, page, limit int) (*querybuilder.PaginateResult, error) {
	ctx := context.Background()
	users := []user.User{}

	opts := querybuilder.PaginateOptions{
		Page:  page,
		Limit: limit,
	}

	result, err := repo.FindWithPagination(ctx, &user.User{}, &users, filter, opts)
	if err != nil {
		return nil, err
	}

	return result, nil
}
