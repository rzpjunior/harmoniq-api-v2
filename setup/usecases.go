package setup

import (
	albumUseCase "harmoniq/harmoniq-api-v2/service/album/usecase"
	artistUsecase "harmoniq/harmoniq-api-v2/service/artist/usecase"
	cartUseCase "harmoniq/harmoniq-api-v2/service/cart/usecase"
	categoryUsecase "harmoniq/harmoniq-api-v2/service/category/usecase"
	productUseCase "harmoniq/harmoniq-api-v2/service/product/usecase"
	userUseCase "harmoniq/harmoniq-api-v2/service/user/usecase"

	"harmoniq/harmoniq-api-v2/domain"

	"time"
)

type UseCases struct {
	ProductUsecase  domain.ProductUsecase
	UserUsecase     domain.UserUsecase
	CategoryUsecase domain.CategoryUsecase
	CartUsecase     domain.CartUsecase
	AlbumUsecase    domain.AlbumUsecase
	ArtistUsecase   domain.ArtistUsecase
}

func NewUseCases(repos *Repositories, timeout time.Duration) *UseCases {
	return &UseCases{
		ProductUsecase:  productUseCase.NewProductUsecase(repos.ProductRepo, repos.CategoryRepo, repos.ProductImageRepo, timeout),
		UserUsecase:     userUseCase.NewUserUsecase(repos.UserRepo, timeout),
		CategoryUsecase: categoryUsecase.NewCategoryUsecase(repos.CategoryRepo, timeout),
		CartUsecase:     cartUseCase.NewCartUsecase(repos.CartRepo, repos.ProductRepo, repos.CategoryRepo, repos.ProductImageRepo, timeout),
		AlbumUsecase:    albumUseCase.NewAlbumUsecase(repos.AlbumRepo, repos.ArtistRepo, timeout),
		ArtistUsecase:   artistUsecase.NewArtistUsecase(repos.ArtistRepo, timeout),
	}
}
