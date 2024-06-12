package setup

import (
	cartHandler "harmoniq/harmoniq-api-v2/service/cart/delivery/http"
	categoryHandler "harmoniq/harmoniq-api-v2/service/category/delivery/http"
	productHandler "harmoniq/harmoniq-api-v2/service/product/delivery/http"
	userHandler "harmoniq/harmoniq-api-v2/service/user/delivery/http"

	"github.com/labstack/echo"
)

func SetupHandlers(e *echo.Echo, useCases *UseCases) {
	productHandler.NewProductHandler(e, useCases.ProductUsecase)
	userHandler.NewUserHandler(e, useCases.UserUsecase)
	categoryHandler.NewCategoryHandler(e, useCases.CategoryUsecase)
	cartHandler.NewCartHandler(e, useCases.CartUsecase)
}
