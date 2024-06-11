package http

import (
	"project-version3/superindo-task/middleware"
	"project-version3/superindo-task/pkg/ehttp"
	"project-version3/superindo-task/service/domain"
	"project-version3/superindo-task/service/domain/dto"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// CartHandler  represent the httphandler for Cart
type CartHandler struct {
	PUsecase domain.CartUsecase
}

// NewCartHandler will initialize the Cart resources endpoint
func NewCartHandler(e *echo.Echo, ps domain.CartUsecase) {
	handler := &CartHandler{
		PUsecase: ps,
	}
	v1 := e.Group("v1")

	mw := middleware.NewMiddleware()

	v1.GET("/cart", handler.GetList, mw.Authorized())
	v1.PUT("/cart", handler.Update, mw.Authorized())
}

func (h CartHandler) GetList(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var carts []dto.CartResponse
	var total int64
	carts, total, err = h.PUsecase.GetList(ctx.Request().Context())
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.DataList(carts, total, 0, 0)

	return ctx.Serve(err)
}

func (h CartHandler) Update(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)
	validator := validator.New()

	var req dto.CartRequestUpdate

	if err = ctx.Bind(&req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	if err = validator.Struct(req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	err = h.PUsecase.Update(ctx.Request().Context(), req)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	return ctx.Serve(err)
}
