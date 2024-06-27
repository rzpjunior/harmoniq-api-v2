package dto

import "time"

type SongResponse struct {
	SongId       int            `json:"song_id"`
	Title        string         `json:"title"`
	Duration     string         `json:"duration"`
	TrackNumber  int            `json:"track_number"`
	AudioFileUrl string         `json:"audio_file_url"`
	ImageUrl     string         `json:"image_url"`
	AlbumId      AlbumResponse  `json:"album,omitempty"`
	ArtistId     ArtistResponse `json:"artist,omitempty"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
}
