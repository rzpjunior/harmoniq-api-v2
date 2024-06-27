package setup

import (
	albumUseCase "harmoniq/harmoniq-api-v2/service/album/usecase"
	artistUsecase "harmoniq/harmoniq-api-v2/service/artist/usecase"
	songUseCase "harmoniq/harmoniq-api-v2/service/song/usecase"
	userUseCase "harmoniq/harmoniq-api-v2/service/user/usecase"

	"harmoniq/harmoniq-api-v2/domain"

	"time"
)

type UseCases struct {
	UserUsecase   domain.UserUsecase
	AlbumUsecase  domain.AlbumUsecase
	ArtistUsecase domain.ArtistUsecase
	SongUsecase   domain.SongUsecase
}

func NewUseCases(repos *Repositories, timeout time.Duration) *UseCases {
	return &UseCases{
		UserUsecase:   userUseCase.NewUserUsecase(repos.UserRepo, timeout),
		AlbumUsecase:  albumUseCase.NewAlbumUsecase(repos.AlbumRepo, repos.ArtistRepo, timeout),
		ArtistUsecase: artistUsecase.NewArtistUsecase(repos.ArtistRepo, timeout),
		SongUsecase:   songUseCase.NewSongUsecase(repos.SongRepo, repos.AlbumRepo, repos.ArtistRepo, timeout),
	}
}
