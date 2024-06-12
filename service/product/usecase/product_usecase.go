package usecase

import (
	"context"
	"time"

	"harmoniq/harmoniq-api-v2/domain"
	"harmoniq/harmoniq-api-v2/domain/dto"

	"github.com/labstack/gommon/log"
)

type productUsecase struct {
	productRepo      domain.ProductRepository
	categoryRepo     domain.CategoryRepository
	productImageRepo domain.ProductImageRepository
	contextTimeout   time.Duration
}

// NewProductUsecase will create new an articleUsecase object representation of domain.ProductUsecase interface
func NewProductUsecase(u domain.ProductRepository, c domain.CategoryRepository, pi domain.ProductImageRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepo:      u,
		categoryRepo:     c,
		productImageRepo: pi,
		contextTimeout:   timeout,
	}
}

func (s *productUsecase) GetList(ctx context.Context, offset int, limit int, search string, categoryId int) (res []dto.ProductResponse, total int64, err error) {
	var products []domain.Product
	products, total, err = s.productRepo.GetList(ctx, offset, limit, search, categoryId)
	if err != nil {
		log.Error(err)
		return
	}

	for _, product := range products {
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

		res = append(res, dto.ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Weight:      product.Weight,
			Price:       product.Price,
			Stock:       product.Stock,
			Status:      product.Status,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			Category: dto.CategoryResponse{
				Id:        category.Id,
				Name:      category.Name,
				IconUrl:   category.IconUrl,
				CreatedAt: category.CreatedAt,
				UpdatedAt: category.UpdatedAt,
			},
			ProductImages: productImageResponse,
		})
	}

	return
}

func (s *productUsecase) GetDetail(ctx context.Context, id int) (res dto.ProductResponse, err error) {
	var product domain.Product
	product, err = s.productRepo.GetDetail(ctx, id)
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

	res = dto.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Weight:      product.Weight,
		Price:       product.Price,
		Stock:       product.Stock,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Category: dto.CategoryResponse{
			Id:        category.Id,
			Name:      category.Name,
			IconUrl:   category.IconUrl,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		},
		ProductImages: productImageResponse,
	}

	return
}
