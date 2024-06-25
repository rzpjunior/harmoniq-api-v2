package setup

import (
	albumRepo "harmoniq/harmoniq-api-v2/service/album/repository/mysql"
	artistRepo "harmoniq/harmoniq-api-v2/service/artist/repository/mysql"
	cartRepo "harmoniq/harmoniq-api-v2/service/cart/repository/mysql"
	categoryRepo "harmoniq/harmoniq-api-v2/service/category/repository/mysql"
	productRepo "harmoniq/harmoniq-api-v2/service/product/repository/mysql"
	productImageRepo "harmoniq/harmoniq-api-v2/service/product_image/repository/mysql"
	userRepo "harmoniq/harmoniq-api-v2/service/user/repository/mysql"

	"harmoniq/harmoniq-api-v2/domain"

	"gorm.io/gorm"
)

type Repositories struct {
	ProductRepo      domain.ProductRepository
	UserRepo         domain.UserRepository
	CategoryRepo     domain.CategoryRepository
	ProductImageRepo domain.ProductImageRepository
	CartRepo         domain.CartRepository
	AlbumRepo        domain.AlbumRepository
	ArtistRepo       domain.ArtistRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ProductRepo:      productRepo.NewMysqlProductRepository(db),
		UserRepo:         userRepo.NewMysqlUserRepository(db),
		CategoryRepo:     categoryRepo.NewMysqlCategoryRepository(db),
		ProductImageRepo: productImageRepo.NewMysqlProductImageRepository(db),
		CartRepo:         cartRepo.NewMysqlCartRepository(db),
		AlbumRepo:        albumRepo.NewMysqlAlbumRepository(db),
		ArtistRepo:       artistRepo.NewMysqlArtistRepository(db),
	}
}
