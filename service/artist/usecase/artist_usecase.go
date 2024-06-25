package usecase

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain"
	"harmoniq/harmoniq-api-v2/domain/dto"
	"time"

	"github.com/labstack/gommon/log"
)

type artistUsecase struct {
	artistRepo     domain.ArtistRepository
	contextTimeout time.Duration
}

func NewArtistUsecase(u domain.ArtistRepository, timeout time.Duration) domain.ArtistUsecase {
	return &artistUsecase{
		artistRepo: u,
	}
}

func (s *artistUsecase) GetList(ctx context.Context, offset int, limit int, search string) (res []dto.ArtistResponse, total int64, err error) {
	var artists []domain.Artist
	artists, total, err = s.artistRepo.GetList(ctx, offset, limit, search)
	if err != nil {
		log.Error(err)
		return
	}

	for _, artist := range artists {
		res = append(res, dto.ArtistResponse{
			Id:   artist.Id,
			Name: artist.Name,
		})
	}

	return
}

func (s *artistUsecase) GetDetail(ctx context.Context, id int) (res dto.ArtistResponse, err error) {
	var artist domain.Artist
	artist, err = s.artistRepo.GetDetail(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	res = dto.ArtistResponse{
		Id:   artist.Id,
		Name: artist.Name,
	}

	return
}
