package dto

type ArtistResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	// RecordLabel RecordLabelResponse `json:"record_label"`
}

type RecordLabelResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
