package usecase

import (
	"context"
	"time"

	"harmoniq/harmoniq-api-v2/pkg/constants"
	"harmoniq/harmoniq-api-v2/pkg/ehttp"
	"harmoniq/harmoniq-api-v2/service/domain"
	"harmoniq/harmoniq-api-v2/service/domain/dto"

	"github.com/labstack/gommon/log"
)

type cartUsecase struct {
	cartRepo         domain.CartRepository
	productRepo      domain.ProductRepository
	categoryRepo     domain.CategoryRepository
	productImageRepo domain.ProductImageRepository
	contextTimeout   time.Duration
}

// NewCartUsecase will create new an articleUsecase object representation of domain.CartUsecase interface
func NewCartUsecase(u domain.CartRepository, p domain.ProductRepository, c domain.CategoryRepository, pi domain.ProductImageRepository, timeout time.Duration) domain.CartUsecase {
	return &cartUsecase{
		cartRepo:         u,
		productRepo:      p,
		categoryRepo:     c,
		productImageRepo: pi,
		contextTimeout:   timeout,
	}
}

func (s *cartUsecase) GetList(ctx context.Context) (res []dto.CartResponse, total int64, err error) {
	userId := ctx.Value(constants.KeyUserID).(int)

	var carts []domain.Cart
	carts, total, err = s.cartRepo.GetListByUserId(ctx, userId)
	if err != nil {
		log.Error(err)
		return
	}

	for _, cart := range carts {
		var product domain.Product
		product, err = s.productRepo.GetDetail(ctx, cart.ProductId)
		if err != nil {
			log.Error(err)
			return
		}

		var category domain.Category
		category, err = s.categoryRepo.GetDetail(ctx, product.CategoryId)
		if err != nil {
			log.Error(err)
			return
		}

		var productImages []domain.ProductImage
		productImages, _, err = s.productImageRepo.GetListByProductId(ctx, product.Id)
		if err != nil {
			log.Error(err)
			return
		}

		var productImageResponse []dto.ProductImageResponse
		for _, productImage := range productImages {
			productImageResponse = append(productImageResponse, dto.ProductImageResponse{
				Id:        productImage.Id,
				ImageUrl:  productImage.ImageUrl,
				MainImage: productImage.MainImage,
			})
		}

		res = append(res, dto.CartResponse{
			Id:  cart.Id,
			Qty: cart.Qty,
			Product: dto.ProductResponse{
				Id:          product.Id,
				Name:        product.Name,
				Description: product.Description,
				Weight:      product.Weight,
				Price:       product.Price,
				Status:      product.Status,
				Stock:       product.Stock,
				Category: dto.CategoryResponse{
					Id:        category.Id,
					Name:      category.Name,
					IconUrl:   category.IconUrl,
					CreatedAt: category.CreatedAt,
					UpdatedAt: category.UpdatedAt,
				},
				ProductImages: productImageResponse,
				CreatedAt:     product.CreatedAt,
				UpdatedAt:     product.UpdatedAt,
			},
		})

	}

	return
}

func (s *cartUsecase) Update(ctx context.Context, req dto.CartRequestUpdate) (err error) {
	userId := ctx.Value(constants.KeyUserID).(int)

	// validate product
	var product domain.Product
	product, err = s.productRepo.GetDetail(ctx, req.ProductId)
	if err != nil {
		log.Error(err)
		err = ehttp.ErrorOutput("product_id", "The product is invalid")
		return
	}

	// validate status
	if product.Status != 1 {
		err = ehttp.ErrorOutput("product_id", "The product must be active")
		return
	}

	// validates stock
	if product.Stock < req.Qty {
		err = ehttp.ErrorOutput("qty", "The quantity should not be more than stock")
		return
	}

	// check active active cart
	var cartActive domain.Cart
	cartActive, _ = s.cartRepo.CheckCart(ctx, userId, req.ProductId)
	if cartActive.Id != 0 {
		cartActive.Qty = req.Qty
		err = s.cartRepo.Update(ctx, &cartActive)
		if err != nil {
			log.Error(err)
			return
		}
		return
	} else {
		cart := &domain.Cart{
			UserId:    userId,
			ProductId: req.ProductId,
			Qty:       req.Qty,
			Status:    1,
		}
		err = s.cartRepo.Create(ctx, cart)
		if err != nil {
			log.Error(err)
			return
		}
	}

	return
}
