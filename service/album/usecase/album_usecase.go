package usecase

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain"
	"harmoniq/harmoniq-api-v2/domain/dto"
	"time"

	"github.com/labstack/gommon/log"
)

type albumUsecase struct {
	albumRepo      domain.AlbumRepository
	artistRepo     domain.ArtistRepository
	contextTimeout time.Duration
}

func NewAlbumUsecase(u domain.AlbumRepository, c domain.ArtistRepository, timeout time.Duration) domain.AlbumUsecase {
	return &albumUsecase{
		albumRepo:      u,
		artistRepo:     c,
		contextTimeout: timeout,
	}
}

func (s *albumUsecase) GetList(ctx context.Context, offset int, limit int, search string, artistId int) (res []dto.AlbumResponse, total int64, err error) {
	var albums []domain.Album
	albums, total, err = s.albumRepo.GetList(ctx, offset, limit, search, artistId)
	if err != nil {
		log.Error(err)
		return
	}

	for _, album := range albums {
		var artist domain.Artist
		artist, err = s.artistRepo.GetDetail(ctx, album.ArtistId)
		if err != nil {
			log.Error(err)
			return
		}

		res = append(res, dto.AlbumResponse{
			Id:   album.Id,
			Name: album.Name,
			Year: album.Year,
			Artist: dto.ArtistResponse{
				Id:   artist.Id,
				Name: artist.Name,
			},
		})
	}

	return
}

func (s *albumUsecase) GetDetail(ctx context.Context, id int) (res dto.AlbumResponse, err error) {
	var album domain.Album
	album, err = s.albumRepo.GetDetail(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	var artist domain.Artist
	artist, err = s.artistRepo.GetDetail(ctx, album.ArtistId)
	if err != nil {
		log.Error(err)
		return
	}

	res = dto.AlbumResponse{
		Id:   album.Id,
		Name: album.Name,
		Year: album.Year,
		Artist: dto.ArtistResponse{
			Id:   artist.Id,
			Name: artist.Name,
		},
	}

	return
}
