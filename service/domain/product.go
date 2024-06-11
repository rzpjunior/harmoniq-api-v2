package domain

import (
	"context"
	"project-version3/superindo-task/service/domain/dto"
	"time"
)

// Product is representing the Product data struct
type Product struct {
	Id          int `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Description string
	CategoryId  int
	Weight      float64
	Price       float64
	Stock       int
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *Product) TableName() string {
	return "product"
}

// ProductUsecase represent the article's usecases
type ProductUsecase interface {
	GetList(ctx context.Context, offset int, limit int, search string, categoryId int) (res []dto.ProductResponse, total int64, err error)
	GetDetail(ctx context.Context, id int) (res dto.ProductResponse, err error)
}

// ProductRepository represent the article's repository contract
type ProductRepository interface {
	GetList(ctx context.Context, offset int, limit int, search string, categoryId int) (products []Product, count int64, err error)
	GetDetail(ctx context.Context, id int) (product Product, err error)
}
