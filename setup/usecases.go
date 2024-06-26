package setup

import (
	albumUseCase "harmoniq/harmoniq-api-v2/service/album/usecase"
	artistUsecase "harmoniq/harmoniq-api-v2/service/artist/usecase"
	userUseCase "harmoniq/harmoniq-api-v2/service/user/usecase"

	"harmoniq/harmoniq-api-v2/domain"

	"time"
)

type UseCases struct {
	UserUsecase   domain.UserUsecase
	AlbumUsecase  domain.AlbumUsecase
	ArtistUsecase domain.ArtistUsecase
}

func NewUseCases(repos *Repositories, timeout time.Duration) *UseCases {
	return &UseCases{
		UserUsecase:   userUseCase.NewUserUsecase(repos.UserRepo, timeout),
		AlbumUsecase:  albumUseCase.NewAlbumUsecase(repos.AlbumRepo, repos.ArtistRepo, timeout),
		ArtistUsecase: artistUsecase.NewArtistUsecase(repos.ArtistRepo, timeout),
	}
}
