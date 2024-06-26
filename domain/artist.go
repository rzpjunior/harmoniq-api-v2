package domain

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain/dto"
	"time"
)

type Artist struct {
	ArtistId  int `gorm:"primaryKey;autoIncrement:true"`
	Name      string
	Genre     string
	Country   string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Artist) TableName() string {
	return "artists"
}

type ArtistUsecase interface {
	GetList(ctx context.Context, offset int, limit int, search string) (res []dto.ArtistResponse, total int64, err error)
	GetDetail(ctx context.Context, id int) (res dto.ArtistResponse, err error)
}

type ArtistRepository interface {
	GetList(ctx context.Context, offset int, limit int, search string) (artists []Artist, count int64, err error)
	GetDetail(ctx context.Context, id int) (artist Artist, err error)
}
