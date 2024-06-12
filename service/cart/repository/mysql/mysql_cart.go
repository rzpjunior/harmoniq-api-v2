package mysql

import (
	"context"
	"harmoniq/harmoniq-api-v2/service/domain"

	"gorm.io/gorm"
)

type mysqlCartRepository struct {
	Conn *gorm.DB
}

func NewMysqlCartRepository(conn *gorm.DB) domain.CartRepository {
	return &mysqlCartRepository{conn}
}

func (m *mysqlCartRepository) GetListByUserId(ctx context.Context, userId int) (carts []domain.Cart, count int64, err error) {
	gorm := m.Conn.Model(domain.Cart{})
	if userId != 0 {
		gorm = gorm.Where("user_id = ?", userId)
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Find(&carts).Error
	return
}

func (m *mysqlCartRepository) CheckCart(ctx context.Context, userId int, productId int) (cart domain.Cart, err error) {
	err = m.Conn.Where("user_id = ?", userId).Where("product_id = ?", productId).Where("status = ?", 1).First(&cart).Error
	return
}

func (m *mysqlCartRepository) Create(ctx context.Context, cart *domain.Cart) (err error) {
	err = m.Conn.Create(cart).Error
	return
}

func (m *mysqlCartRepository) Update(ctx context.Context, cart *domain.Cart) (err error) {
	err = m.Conn.Save(cart).Error
	return
}
