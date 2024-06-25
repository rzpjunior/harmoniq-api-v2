package http

import (
	"harmoniq/harmoniq-api-v2/domain"
	"harmoniq/harmoniq-api-v2/domain/dto"
	"harmoniq/harmoniq-api-v2/middleware"
	"harmoniq/harmoniq-api-v2/pkg/ehttp"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type AlbumHandler struct {
	AUsecase domain.AlbumUsecase
}

func NewAlbumHandler(e *echo.Echo, ps domain.AlbumUsecase) {
	handler := &AlbumHandler{
		AUsecase: ps,
	}
	v1 := e.Group("v1")

	mw := middleware.NewMiddleware()

	v1.GET("/album", handler.GetList, mw.Authorized())
	v1.GET("/album/:id", handler.GetDetail, mw.Authorized())
}

func (h AlbumHandler) GetList(c echo.Context) (err error) {
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
	artistId := ctx.GetParamInt("artist_id")

	var albums []dto.AlbumResponse
	var total int64
	albums, total, err = h.AUsecase.GetList(ctx.Request().Context(), page.Start, page.Limit, search, artistId)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.DataList(albums, total, page.Page, page.PerPage)

	return ctx.Serve(err)
}

func (h AlbumHandler) GetDetail(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var id int
	id, err = ctx.GetParamUri("id")
	if err != nil {
		log.Error(err)
		return
	}

	var album dto.AlbumResponse
	album, err = h.AUsecase.GetDetail(ctx.Request().Context(), id)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.Data(album)

	return ctx.Serve(err)
}
