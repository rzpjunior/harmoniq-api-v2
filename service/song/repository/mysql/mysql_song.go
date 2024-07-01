package mysql

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain"

	"gorm.io/gorm"
)

type mysqlSongRepository struct {
	Conn *gorm.DB
}

func NewMysqlSongRepository(conn *gorm.DB) domain.SongRepository {
	return &mysqlSongRepository{conn}
}

func (m *mysqlSongRepository) GetList(ctx context.Context, offset int, limit int, search string, albumId int, artistId int) (songs []domain.Song, count int64, err error) {
	gorm := m.Conn.Model(domain.Song{})

	if search != "" {
		gorm = gorm.Where("title like ?", "%"+search+"%")
	}
	if albumId != 0 {
		gorm = gorm.Where("album_id = ?", albumId)
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Offset(offset).Limit(limit).Find(&songs).Error

	return
}

func (m *mysqlSongRepository) GetDetail(ctx context.Context, id int) (song domain.Song, err error) {
	err = m.Conn.Where("song_id = ?", id).First(&song).Error
	return
}
