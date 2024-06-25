package domain

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain/dto"
)

type Album struct {
	Id       int `gorm:"primaryKey;autoIncrement:true"`
	ArtistId int
	Name     string
	Year     int
}

func (m *Album) TableName() string {
	return "album"
}

type AlbumUsecase interface {
	GetList(ctx context.Context, offset int, limit int, search string, artistId int) (res []dto.AlbumResponse, total int64, err error)
	GetDetail(ctx context.Context, id int) (res dto.AlbumResponse, err error)
}

type AlbumRepository interface {
	GetList(ctx context.Context, offset int, limit int, search string, artistId int) (albums []Album, count int64, err error)
	GetDetail(ctx context.Context, id int) (album Album, err error)
}
