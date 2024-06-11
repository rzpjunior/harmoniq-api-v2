package domain

import (
	"context"
	"project-version3/superindo-task/service/domain/dto"
	"time"
)

type Category struct {
	Id        int `gorm:"primaryKey;autoIncrement:true"`
	Name      string
	IconUrl   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Category) TableName() string {
	return "category"
}

// CategoryUsecase represent the article's usecases
type CategoryUsecase interface {
	GetList(ctx context.Context, offset int, limit int, search string) (res []dto.CategoryResponse, total int64, err error)
	GetDetail(ctx context.Context, id int) (res dto.CategoryResponse, err error)
}

// CategoryRepository represent the article's repository contract
type CategoryRepository interface {
	GetList(ctx context.Context, offset int, limit int, search string) (categories []Category, count int64, err error)
	GetDetail(ctx context.Context, id int) (category Category, err error)
}
