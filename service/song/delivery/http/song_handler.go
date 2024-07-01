package http

import (
	"harmoniq/harmoniq-api-v2/domain"
	"harmoniq/harmoniq-api-v2/domain/dto"
	"harmoniq/harmoniq-api-v2/middleware"
	"harmoniq/harmoniq-api-v2/pkg/ehttp"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type SongHandler struct {
	SUsecase domain.SongUsecase
}

func NewSongHandler(e *echo.Echo, ps domain.SongUsecase) {
	handler := &SongHandler{
		SUsecase: ps,
	}
	v1 := e.Group("v1")

	mw := middleware.NewMiddleware()

	v1.GET("/song", handler.GetList, mw.Authorized())
	v1.GET("/song/:id", handler.GetDetail, mw.Authorized())
}

func (h SongHandler) GetList(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	// get pagination
	var page *ehttp.Paginator
	page, err = ehttp.NewPaginator(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	// get params
	search := ctx.GetParamString("search")
	albumId := ctx.GetParamInt("album_id")
	artistId := ctx.GetParamInt("artist_id")

	var songs []dto.SongResponse
	var total int64
	songs, total, err = h.SUsecase.GetList(ctx.Request().Context(), page.Start, page.Limit, search, artistId, albumId)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.DataList(songs, total, page.Page, page.PerPage)

	return ctx.Serve(err)
}

func (h SongHandler) GetDetail(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var id int
	id, err = ctx.GetParamUri("id")
	if err != nil {
		log.Error(err)
		return
	}

	var song dto.SongResponse
	song, err = h.SUsecase.GetDetail(ctx.Request().Context(), id)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.Data(song)

	return ctx.Serve(err)
}
