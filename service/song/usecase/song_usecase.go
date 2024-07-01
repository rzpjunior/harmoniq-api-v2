package usecase

import (
	"context"
	"harmoniq/harmoniq-api-v2/domain"
	"harmoniq/harmoniq-api-v2/domain/dto"
	"time"

	"github.com/labstack/gommon/log"
)

type songUsecase struct {
	songRepo       domain.SongRepository
	artistRepo     domain.ArtistRepository
	albumRepo      domain.AlbumRepository
	artistSongRepo domain.ArtistSongRepository
	contextTimeout time.Duration
}

func NewSongUsecase(u domain.SongRepository, c domain.AlbumRepository, r domain.ArtistRepository, s domain.ArtistSongRepository, timeout time.Duration) domain.SongUsecase {
	return &songUsecase{
		songRepo:       u,
		albumRepo:      c,
		artistRepo:     r,
		artistSongRepo: s,
		contextTimeout: timeout,
	}
}

func (s *songUsecase) GetList(ctx context.Context, offset int, limit int, search string, albumId int, artistId int) (res []dto.SongResponse, total int64, err error) {
	var songs []domain.Song
	songs, total, err = s.songRepo.GetList(ctx, offset, limit, search, albumId, artistId)
	if err != nil {
		log.Error(err)
		return
	}

	for _, song := range songs {
		var album domain.Album
		if song.AlbumId != 0 {
			album, err = s.albumRepo.GetDetail(ctx, song.AlbumId)
			if err != nil {
				log.Error(err)
				return
			}
		}

		var artist domain.Artist
		if album.ArtistId != 0 {
			artist, err = s.artistRepo.GetDetail(ctx, album.ArtistId)
			if err != nil {
				log.Error(err)
				return
			}
		}

		dtoSongResponse := dto.SongResponse{
			SongId:       song.SongId,
			Title:        song.Title,
			Duration:     song.Duration,
			TrackNumber:  song.TrackNumber,
			AudioFileUrl: song.AudioFileUrl,
			CreatedAt:    song.CreatedAt,
			UpdatedAt:    song.UpdatedAt,
		}

		if song.AlbumId != 0 {
			dtoSongResponse.AlbumId = &dto.AlbumResponse{
				AlbumId:       album.AlbumId,
				Title:         album.Title,
				Genre:         album.Genre,
				ReleaseDate:   album.ReleaseDate,
				CoverImageUrl: album.CoverImageUrl,
				Artist: dto.ArtistResponse{
					ArtistId: artist.ArtistId,
					Name:     artist.Name,
					Bio:      artist.Bio,
					Country:  artist.Country,
					Genre:    artist.Genre,
				},
			}
		} else {
			var artistSong []domain.ArtistSong
			artistSong, _, err = s.artistSongRepo.GetList(ctx, 0, 1000, "", 0, song.SongId)
			if err != nil {
				log.Error(err)
				return
			}
			if len(artistSong) > 0 {
				var artist1 domain.Artist
				artist1, err = s.artistRepo.GetDetail(ctx, artistSong[0].ArtistId)
				if err != nil {
					log.Error(err)
					return
				}
				dtoSongResponse.ArtistId = &dto.ArtistResponse{
					ArtistId: artist1.ArtistId,
					Name:     artist1.Name,
					Bio:      artist1.Bio,
					Country:  artist1.Country,
					Genre:    artist1.Genre,
				}
			}
		}

		res = append(res, dtoSongResponse)
	}

	return
}

func (s *songUsecase) GetDetail(ctx context.Context, id int) (res dto.SongResponse, err error) {
	var song domain.Song
	song, err = s.songRepo.GetDetail(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	var album domain.Album
	album, err = s.albumRepo.GetDetail(ctx, album.AlbumId)
	if err != nil {
		log.Error(err)
		return
	}

	var artist domain.Artist
	artist, err = s.artistRepo.GetDetail(ctx, artist.ArtistId)
	if err != nil {
		log.Error(err)
		return
	}

	res = dto.SongResponse{
		SongId:       song.SongId,
		Title:        song.Title,
		Duration:     song.Duration,
		TrackNumber:  song.TrackNumber,
		AudioFileUrl: song.AudioFileUrl,
		AlbumId: &dto.AlbumResponse{
			AlbumId:       album.AlbumId,
			Title:         album.Title,
			Genre:         album.Genre,
			ReleaseDate:   album.ReleaseDate,
			CoverImageUrl: album.CoverImageUrl,
			Artist: dto.ArtistResponse{
				ArtistId: artist.ArtistId,
				Name:     artist.Name,
				Bio:      artist.Bio,
				Country:  artist.Country,
				Genre:    artist.Genre,
			},
		},
		CreatedAt: song.CreatedAt,
		UpdatedAt: song.UpdatedAt,
	}

	return
}
