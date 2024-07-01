package setup

import (
	albumRepo "harmoniq/harmoniq-api-v2/service/album/repository/mysql"
	artistRepo "harmoniq/harmoniq-api-v2/service/artist/repository/mysql"
	artistSongRepo "harmoniq/harmoniq-api-v2/service/artist_song/repository/mysql"
	songRepo "harmoniq/harmoniq-api-v2/service/song/repository/mysql"
	userRepo "harmoniq/harmoniq-api-v2/service/user/repository/mysql"

	"harmoniq/harmoniq-api-v2/domain"

	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo       domain.UserRepository
	AlbumRepo      domain.AlbumRepository
	ArtistRepo     domain.ArtistRepository
	ArtistSongRepo domain.ArtistSongRepository
	SongRepo       domain.SongRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo:       userRepo.NewMysqlUserRepository(db),
		AlbumRepo:      albumRepo.NewMysqlAlbumRepository(db),
		ArtistRepo:     artistRepo.NewMysqlArtistRepository(db),
		ArtistSongRepo: artistSongRepo.NewMysqlArtistSongRepository(db),
		SongRepo:       songRepo.NewMysqlSongRepository(db),
	}
}
