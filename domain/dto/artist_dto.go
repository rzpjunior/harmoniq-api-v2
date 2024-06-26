package dto

import "time"

type ArtistResponse struct {
	ArtistId  int       `json:"artist_id"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	Country   string    `json:"country"`
	Genre     string    `json:"genre"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
