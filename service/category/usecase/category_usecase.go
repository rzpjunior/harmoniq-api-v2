package usecase

import (
	"context"
	"time"

	"harmoniq/harmoniq-api-v2/service/domain"
	"harmoniq/harmoniq-api-v2/service/domain/dto"

	"github.com/labstack/gommon/log"
)

type categoryUsecase struct {
	categoryRepo domain.CategoryRepository
	// authorRepo     domain.AuthorRepository
	contextTimeout time.Duration
}

// NewCategoryUsecase will create new an articleUsecase object representation of domain.CategoryUsecase interface
func NewCategoryUsecase(u domain.CategoryRepository, timeout time.Duration) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepo:   u,
		contextTimeout: timeout,
	}
}

func (s *categoryUsecase) GetList(ctx context.Context, offset int, limit int, search string) (res []dto.CategoryResponse, total int64, err error) {
	var categories []domain.Category
	categories, total, err = s.categoryRepo.GetList(ctx, offset, limit, search)
	if err != nil {
		log.Error(err)
		return
	}

	for _, category := range categories {
		res = append(res, dto.CategoryResponse{
			Id:        category.Id,
			Name:      category.Name,
			IconUrl:   category.IconUrl,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		})
	}

	return
}

func (s *categoryUsecase) GetDetail(ctx context.Context, id int) (res dto.CategoryResponse, err error) {
	var category domain.Category
	category, err = s.categoryRepo.GetDetail(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	res = dto.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		IconUrl:   category.IconUrl,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}

	return
}
