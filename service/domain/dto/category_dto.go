package dto

import "time"

type CategoryResponse struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	IconUrl   string            `json:"icon_url"`
	Products  []ProductResponse `json:"products,omitempty"`
	CreatedAt time.Time         `json:"created_at,omitempty"`
	UpdatedAt time.Time         `json:"updated_at,omitempty"`
}
