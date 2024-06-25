package mysql

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain"

	"gorm.io/gorm"
)

type mysqlArtistRepository struct {
	Conn *gorm.DB
}

func NewMysqlArtistRepository(conn *gorm.DB) domain.ArtistRepository {
	return &mysqlArtistRepository{conn}
}

func (m *mysqlArtistRepository) GetList(ctx context.Context, offset int, limit int, search string) (artists []domain.Artist, count int64, err error) {
	gorm := m.Conn.Model(domain.Album{})

	if search != "" {
		gorm = gorm.Where("name like ?", "%"+search+"%")
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Offset(offset).Limit(limit).Find(&artists).Error

	return
}

func (m *mysqlArtistRepository) GetDetail(ctx context.Context, id int) (artist domain.Artist, err error) {
	err = m.Conn.Where("id = ?", id).First(&artist).Error
	return
}
