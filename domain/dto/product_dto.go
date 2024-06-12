package dto

import "time"

type ProductResponse struct {
	Id            int                    `json:"id"`
	Name          string                 `json:"name"`
	Description   string                 `json:"description"`
	Weight        float64                `json:"weight"`
	Price         float64                `json:"price"`
	Status        int                    `json:"status"`
	Stock         int                    `json:"stock"`
	Category      CategoryResponse       `json:"category,omitempty"`
	ProductImages []ProductImageResponse `json:"product_images,omitempty"`
	CreatedAt     time.Time              `json:"created_at,omitempty"`
	UpdatedAt     time.Time              `json:"updated_at,omitempty"`
}

type ProductImageResponse struct {
	Id        int       `json:"id"`
	ProductId int       `json:"product_id,omitempty"`
	ImageUrl  string    `json:"image_url,omitempty"`
	MainImage int       `json:"main_image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
