package mysql

import (
	"context"
	"harmoniq/harmoniq-api-v2/service/domain"

	"gorm.io/gorm"
)

type mysqlCategoryRepository struct {
	Conn *gorm.DB
}

func NewMysqlCategoryRepository(conn *gorm.DB) domain.CategoryRepository {
	return &mysqlCategoryRepository{conn}
}

func (m *mysqlCategoryRepository) GetList(ctx context.Context, offset int, limit int, search string) (categories []domain.Category, count int64, err error) {
	gorm := m.Conn.Model(domain.Category{})
	if search != "" {
		gorm = gorm.Where("name like ?", "%"+search+"%")
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Offset(offset).Limit(limit).Find(&categories).Error

	return
}

func (m *mysqlCategoryRepository) GetDetail(ctx context.Context, id int) (category domain.Category, err error) {
	err = m.Conn.Where("id = ?", id).First(&category).Error
	return
}
