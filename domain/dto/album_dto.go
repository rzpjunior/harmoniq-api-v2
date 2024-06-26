package dto

import "time"

type AlbumResponse struct {
	AlbumId       int            `json:"album_id"`
	Title         string         `json:"title"`
	Genre         string         `json:"genre"`
	ReleaseDate   string         `json:"release_date"`
	CoverImageUrl string         `json:"cover_image_url"`
	Artist        ArtistResponse `json:"artist"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty"`
}
