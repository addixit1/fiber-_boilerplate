package userv1

import "time"

// CreateUserDTO for creating a new user
type CreateUserDTO struct {
	Name  string `json:"name" validate:"required" example:"Aman"`
	Email string `json:"email" validate:"required,email" example:"aman@gmail.com"`
}

// UserResponseDTO for API responses (Swagger compatible)
type UserResponseDTO struct {
	ID        string    `json:"id" example:"507f1f77bcf86cd799439011"`
	Name      string    `json:"name" example:"Aman"`
	Email     string    `json:"email" example:"aman@gmail.com"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T12:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-01T12:00:00Z"`
}
