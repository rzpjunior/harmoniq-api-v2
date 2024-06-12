package domain

import (
	"context"
	"harmoniq/harmoniq-api-v2/service/domain/dto"
	"time"
)

type Cart struct {
	Id        int `gorm:"primaryKey;autoIncrement:true"`
	UserId    int
	ProductId int
	Qty       int
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Cart) TableName() string {
	return "cart"
}

// CartUsecase represent the article's usecases
type CartUsecase interface {
	GetList(ctx context.Context) (res []dto.CartResponse, total int64, err error)
	Update(ctx context.Context, req dto.CartRequestUpdate) (err error)
}

// CartRepository represent the article's repository contract
type CartRepository interface {
	GetListByUserId(ctx context.Context, userId int) (carts []Cart, count int64, err error)
	CheckCart(ctx context.Context, userId int, productId int) (cart Cart, err error)
	Create(ctx context.Context, cart *Cart) (err error)
	Update(ctx context.Context, cart *Cart) (err error)
}
