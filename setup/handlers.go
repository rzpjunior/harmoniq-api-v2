package setup

import (
	albumHandler "harmoniq/harmoniq-api-v2/service/album/delivery/http"
	artistHandler "harmoniq/harmoniq-api-v2/service/artist/delivery/http"
	userHandler "harmoniq/harmoniq-api-v2/service/user/delivery/http"

	"github.com/labstack/echo"
)

func SetupHandlers(e *echo.Echo, useCases *UseCases) {
	userHandler.NewUserHandler(e, useCases.UserUsecase)
	albumHandler.NewAlbumHandler(e, useCases.AlbumUsecase)
	artistHandler.NewArtistHandler(e, useCases.ArtistUsecase)
}
