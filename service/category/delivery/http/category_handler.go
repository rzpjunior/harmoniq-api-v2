package http

import (
	"harmoniq/harmoniq-api-v2/domain"
	"harmoniq/harmoniq-api-v2/domain/dto"
	"harmoniq/harmoniq-api-v2/middleware"
	"harmoniq/harmoniq-api-v2/pkg/ehttp"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// CategoryHandler  represent the httphandler for Category
type CategoryHandler struct {
	PUsecase domain.CategoryUsecase
}

// NewCategoryHandler will initialize the Category resources endpoint
func NewCategoryHandler(e *echo.Echo, ps domain.CategoryUsecase) {
	handler := &CategoryHandler{
		PUsecase: ps,
	}
	v1 := e.Group("v1")

	mw := middleware.NewMiddleware()

	v1.GET("/category", handler.GetList, mw.Authorized())
	v1.GET("/category/:id", handler.GetDetail, mw.Authorized())
}

func (h CategoryHandler) GetList(c echo.Context) (err error) {
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

	var categories []dto.CategoryResponse
	var total int64
	categories, total, err = h.PUsecase.GetList(ctx.Request().Context(), page.Start, page.Limit, search)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.DataList(categories, total, page.Page, page.PerPage)

	return ctx.Serve(err)
}

func (h CategoryHandler) GetDetail(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var id int
	id, err = ctx.GetParamUri("id")
	if err != nil {
		log.Error(err)
		return
	}

	var category dto.CategoryResponse
	category, err = h.PUsecase.GetDetail(ctx.Request().Context(), id)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.Data(category)

	return ctx.Serve(err)
}
