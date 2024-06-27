package domain

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain/dto"
	"time"
)

type Song struct {
	SongId       int `gorm:"primaryKey;autoIncrement:true"`
	Title        string
	Duration     string
	TrackNumber  int
	AudioFileUrl string
	ImageUrl     string
	AlbumId      int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m *Song) TableName() string {
	return "songs"
}

type SongUsecase interface {
	GetList(ctx context.Context, offset int, limit int, search string, artistId int, albumId int) (res []dto.SongResponse, total int64, err error)
	GetDetail(ctx context.Context, id int) (res dto.SongResponse, err error)
}

type SongRepository interface {
	GetList(ctx context.Context, offset int, limit int, search string, artistId int, albumId int) (songs []Song, count int64, err error)
	GetDetail(ctx context.Context, id int) (song Song, err error)
}
