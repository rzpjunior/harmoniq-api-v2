package mysql

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain"

	"gorm.io/gorm"
)

type mysqlArtistSongRepository struct {
	Conn *gorm.DB
}

func NewMysqlArtistSongRepository(conn *gorm.DB) domain.ArtistSongRepository {
	return &mysqlArtistSongRepository{conn}
}

func (m *mysqlArtistSongRepository) GetList(ctx context.Context, offset int, limit int, search string, songId int, artistId int) (songs []domain.ArtistSong, count int64, err error) {
	gorm := m.Conn.Model(domain.ArtistSong{})

	if songId != 0 {
		gorm = gorm.Where("song_id = ?", songId)
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Offset(offset).Limit(limit).Find(&songs).Error

	return
}

func (m *mysqlArtistSongRepository) GetDetail(ctx context.Context, id int) (song domain.ArtistSong, err error) {
	err = m.Conn.Where("song_id = ?", id).First(&song).Error
	return
}
