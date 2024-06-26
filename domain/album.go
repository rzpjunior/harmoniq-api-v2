package domain

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain/dto"
	"time"
)

type Album struct {
	AlbumId       int `gorm:"primaryKey;autoIncrement:true"`
	ArtistId      int
	Title         string
	ReleaseDate   string
	Genre         string
	CoverImageUrl string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (m *Album) TableName() string {
	return "albums"
}

type AlbumUsecase interface {
	GetList(ctx context.Context, offset int, limit int, search string, artistId int) (res []dto.AlbumResponse, total int64, err error)
	GetDetail(ctx context.Context, id int) (res dto.AlbumResponse, err error)
}

type AlbumRepository interface {
	GetList(ctx context.Context, offset int, limit int, search string, artistId int) (albums []Album, count int64, err error)
	GetDetail(ctx context.Context, id int) (album Album, err error)
}
