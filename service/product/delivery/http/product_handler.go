package http

import (
	"harmoniq/harmoniq-api-v2/middleware"
	"harmoniq/harmoniq-api-v2/pkg/ehttp"
	"harmoniq/harmoniq-api-v2/service/domain"
	"harmoniq/harmoniq-api-v2/service/domain/dto"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// ProductHandler  represent the httphandler for Product
type ProductHandler struct {
	PUsecase domain.ProductUsecase
}

// NewProductHandler will initialize the Product resources endpoint
func NewProductHandler(e *echo.Echo, ps domain.ProductUsecase) {
	handler := &ProductHandler{
		PUsecase: ps,
	}
	v1 := e.Group("v1")

	mw := middleware.NewMiddleware()

	v1.GET("/product", handler.GetList, mw.Authorized())
	v1.GET("/product/:id", handler.GetDetail, mw.Authorized())
}

func (h ProductHandler) GetList(c echo.Context) (err error) {
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
	categoryId := ctx.GetParamInt("category_id")

	var products []dto.ProductResponse
	var total int64
	products, total, err = h.PUsecase.GetList(ctx.Request().Context(), page.Start, page.Limit, search, categoryId)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.DataList(products, total, page.Page, page.PerPage)

	return ctx.Serve(err)
}

func (h ProductHandler) GetDetail(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var id int
	id, err = ctx.GetParamUri("id")
	if err != nil {
		log.Error(err)
		return
	}

	var product dto.ProductResponse
	product, err = h.PUsecase.GetDetail(ctx.Request().Context(), id)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.Data(product)

	return ctx.Serve(err)
}
