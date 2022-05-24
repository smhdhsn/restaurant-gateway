package model

import (
	"time"
)

// UserDTO represents user's data transfer object.
type UserDTO struct {
	ID        uint32
	FirstName string
	LastName  string
	Email     string
	Status    string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// userResp is the response schema.
type userResp struct {
	ID        uint32    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToResp creates a response from dto.
func (s *UserDTO) ToResp() userResp {
	return userResp{
		ID:        s.ID,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Email:     s.Email,
		Status:    s.Status,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}
