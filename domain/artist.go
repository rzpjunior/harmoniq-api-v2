package domain

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain/dto"
)

type Artist struct {
	Id int `gorm:"primaryKey;autoIncrement:true"`
	// RecordLabelId int
	Name string
}

func (m *Artist) TableName() string {
	return "artist"
}

type ArtistUsecase interface {
	GetList(ctx context.Context, offset int, limit int, search string) (res []dto.ArtistResponse, total int64, err error)
	GetDetail(ctx context.Context, id int) (res dto.ArtistResponse, err error)
}

type ArtistRepository interface {
	GetList(ctx context.Context, offset int, limit int, search string) (artists []Artist, count int64, err error)
	GetDetail(ctx context.Context, id int) (artist Artist, err error)
}
