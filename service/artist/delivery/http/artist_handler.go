package http

import (
	"harmoniq/harmoniq-api-v2/domain"
	"harmoniq/harmoniq-api-v2/domain/dto"
	"harmoniq/harmoniq-api-v2/middleware"
	"harmoniq/harmoniq-api-v2/pkg/ehttp"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type ArtistHandler struct {
	ARUsecase domain.ArtistUsecase
}

func NewArtistHandler(e *echo.Echo, ps domain.ArtistUsecase) {
	handler := &ArtistHandler{
		ARUsecase: ps,
	}
	v1 := e.Group("v1")

	mw := middleware.NewMiddleware()

	v1.GET("/artist", handler.GetList, mw.Authorized())
	v1.GET("/artist/:id", handler.GetDetail, mw.Authorized())
}

func (h ArtistHandler) GetList(c echo.Context) (err error) {
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

	var artists []dto.ArtistResponse
	var total int64
	artists, total, err = h.ARUsecase.GetList(ctx.Request().Context(), page.Start, page.Limit, search)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.DataList(artists, total, page.Page, page.PerPage)

	return ctx.Serve(err)
}

func (h ArtistHandler) GetDetail(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var id int
	id, err = ctx.GetParamUri("id")
	if err != nil {
		log.Error(err)
		return
	}

	var artist dto.ArtistResponse
	artist, err = h.ARUsecase.GetDetail(ctx.Request().Context(), id)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.Data(artist)

	return ctx.Serve(err)
}
