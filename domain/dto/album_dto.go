package dto

type AlbumResponse struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Year   int            `json:"year"`
	Artist ArtistResponse `json:"artist"`
}
