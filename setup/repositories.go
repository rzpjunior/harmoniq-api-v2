package setup

import (
	albumRepo "harmoniq/harmoniq-api-v2/service/album/repository/mysql"
	artistRepo "harmoniq/harmoniq-api-v2/service/artist/repository/mysql"
	userRepo "harmoniq/harmoniq-api-v2/service/user/repository/mysql"

	"harmoniq/harmoniq-api-v2/domain"

	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo   domain.UserRepository
	AlbumRepo  domain.AlbumRepository
	ArtistRepo domain.ArtistRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo:   userRepo.NewMysqlUserRepository(db),
		AlbumRepo:  albumRepo.NewMysqlAlbumRepository(db),
		ArtistRepo: artistRepo.NewMysqlArtistRepository(db),
	}
}
