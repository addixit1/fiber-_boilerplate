package user

import (
	"github.com/addixit1/fiber-boilerplate/internal/config"
	"github.com/kamva/mgm/v3"
)

// User model with MGM integration
type User struct {
	// MGM's DefaultModel includes: ID, CreatedAt, UpdatedAt
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name" json:"name"`
	Email            string `bson:"email" json:"email"`
}

// CollectionName returns the MongoDB collection name for User model
func (User) CollectionName() string {
	return config.USERS_COLLECTION
}
