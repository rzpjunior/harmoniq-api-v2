package dto

import "time"

type UserResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	LastLoginAt time.Time `json:"last_login_at,omitempty"`
}
