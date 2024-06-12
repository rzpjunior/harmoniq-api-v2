package domain

import (
	"context"
	"time"
)

// ProductImage is representing the ProductImage data struct
type ProductImage struct {
	Id        int `gorm:"primaryKey;autoIncrement:true"`
	ProductId int
	ImageUrl  string
	MainImage int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *ProductImage) TableName() string {
	return "product_image"
}

// ProductUsecase represent the article's usecases
type ProductImageUsecase interface {
}

// ProductRepository represent the article's repository contract
type ProductImageRepository interface {
	GetListByProductId(ctx context.Context, productId int) (productImages []ProductImage, count int64, err error)
	GetDetail(ctx context.Context, id int) (productImage ProductImage, err error)
}
