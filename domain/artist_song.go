package domain

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

// type SongRepository interface {
// 	GetList(ctx context.Context, offset int, limit int, search string, artistId int, albumId int) (songs []Song, count int64, err error)
// 	GetDetail(ctx context.Context, id int) (song Song, err error)
// }
