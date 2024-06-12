package mysql

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain"

	"gorm.io/gorm"
)

type mysqlProductImageRepository struct {
	Conn *gorm.DB
}

func NewMysqlProductImageRepository(conn *gorm.DB) domain.ProductImageRepository {
	return &mysqlProductImageRepository{conn}
}

func (m *mysqlProductImageRepository) GetListByProductId(ctx context.Context, productId int) (productImages []domain.ProductImage, count int64, err error) {
	gorm := m.Conn.Model(domain.ProductImage{})
	if productId != 0 {
		gorm = gorm.Where("product_id = ?", productId)
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Find(&productImages).Error

	return
}

func (m *mysqlProductImageRepository) GetDetail(ctx context.Context, id int) (productImage domain.ProductImage, err error) {
	err = m.Conn.Where("id = ?", id).First(&productImage).Error
	return
}
