package mysql

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain"

	"gorm.io/gorm"
)

type mysqlAlbumRepository struct {
	Conn *gorm.DB
}

func NewMysqlAlbumRepository(conn *gorm.DB) domain.AlbumRepository {
	return &mysqlAlbumRepository{conn}
}

func (m *mysqlAlbumRepository) GetList(ctx context.Context, offset int, limit int, search string, artistId int) (albums []domain.Album, count int64, err error) {
	gorm := m.Conn.Model(domain.Album{})

	if search != "" {
		gorm = gorm.Where("title like ?", "%"+search+"%")
	}
	if artistId != 0 {
		gorm = gorm.Where("artist_id = ?", artistId)
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Offset(offset).Limit(limit).Find(&albums).Error

	return
}

func (m *mysqlAlbumRepository) GetDetail(ctx context.Context, id int) (album domain.Album, err error) {
	err = m.Conn.Where("album_id = ?", id).First(&album).Error
	return
}
