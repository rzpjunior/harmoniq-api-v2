package domain

import "context"

type ArtistSong struct {
	ArtistId int `gorm:"foreignKey:ArtistId"`
	SongId   int `gorm:"foreignKey:SongId"`
}

func (m *ArtistSong) TableName() string {
	return "artist_song"
}

// type SongUsecase interface {
// 	GetList(ctx context.Context, offset int, limit int, search string, artistId int, albumId int) (res []dto.SongResponse, total int64, err error)
// 	GetDetail(ctx context.Context, id int) (res dto.SongResponse, err error)
// }

type ArtistSongRepository interface {
	GetList(ctx context.Context, offset int, limit int, search string, artistId int, songId int) (songs []ArtistSong, count int64, err error)
}
