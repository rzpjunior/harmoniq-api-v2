package mysql

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	Conn *gorm.DB
}

func NewMysqlProductRepository(conn *gorm.DB) domain.ProductRepository {
	return &mysqlProductRepository{conn}
}
func (m *mysqlProductRepository) GetList(ctx context.Context, offset int, limit int, search string, categoryId int) (products []domain.Product, count int64, err error) {
	gorm := m.Conn.Model(domain.Product{})

	if search != "" {
		gorm = gorm.Where("name like ?", "%"+search+"%")
	}
	if categoryId != 0 {
		gorm = gorm.Where("category_id = ?", categoryId)
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Offset(offset).Limit(limit).Find(&products).Error

	return
}

func (m *mysqlProductRepository) GetDetail(ctx context.Context, id int) (product domain.Product, err error) {
	err = m.Conn.Where("id = ?", id).First(&product).Error
	return
}
