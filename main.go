package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"harmoniq/harmoniq-api-v2/pkg/ehttp"
	_cartHandler "harmoniq/harmoniq-api-v2/service/cart/delivery/http"
	_cartRepo "harmoniq/harmoniq-api-v2/service/cart/repository/mysql"
	_cartUseCase "harmoniq/harmoniq-api-v2/service/cart/usecase"
	_categoryHandler "harmoniq/harmoniq-api-v2/service/category/delivery/http"
	_categoryRepo "harmoniq/harmoniq-api-v2/service/category/repository/mysql"
	_categoryUseCase "harmoniq/harmoniq-api-v2/service/category/usecase"
	_productHandler "harmoniq/harmoniq-api-v2/service/product/delivery/http"
	_productRepo "harmoniq/harmoniq-api-v2/service/product/repository/mysql"
	_productUseCase "harmoniq/harmoniq-api-v2/service/product/usecase"
	_productImageRepo "harmoniq/harmoniq-api-v2/service/product_image/repository/mysql"
	_userHandler "harmoniq/harmoniq-api-v2/service/user/delivery/http"
	_userRepo "harmoniq/harmoniq-api-v2/service/user/repository/mysql"
	_userUseCase "harmoniq/harmoniq-api-v2/service/user/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := gormDB.DB()
	defer func() {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// setup machine and middleware
	e := echo.New()
	// setup cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header"},
	}))

	// setup echo for request id
	e.Use(middleware.RequestID())

	// setup echo for secure
	e.Use(middleware.Secure())

	// setup echo for gzip compres
	e.Use(middleware.Gzip())

	// setup custom context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ehttp.Context{
				Context:        c,
				ResponseFormat: ehttp.NewResponse(),
				ResponseData:   nil,
			}
			return next(cc)
		}
	})

	// setup timeout
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// setup repo
	productRepo := _productRepo.NewMysqlProductRepository(gormDB)
	userRepo := _userRepo.NewMysqlUserRepository(gormDB)
	categoryRepo := _categoryRepo.NewMysqlCategoryRepository(gormDB)
	productImageRepo := _productImageRepo.NewMysqlProductImageRepository(gormDB)
	cartRepo := _cartRepo.NewMysqlCartRepository(gormDB)

	// setup usecase
	productUsecase := _productUseCase.NewProductUsecase(productRepo, categoryRepo, productImageRepo, timeoutContext)
	userUsecase := _userUseCase.NewUserUsecase(userRepo, timeoutContext)
	categoryUsecase := _categoryUseCase.NewCategoryUsecase(categoryRepo, timeoutContext)
	cartUsecase := _cartUseCase.NewCartUsecase(cartRepo, productRepo, categoryRepo, productImageRepo, timeoutContext)

	// setup handler
	_productHandler.NewProductHandler(e, productUsecase)
	_userHandler.NewUserHandler(e, userUsecase)
	_categoryHandler.NewCategoryHandler(e, categoryUsecase)
	_categoryHandler.NewCategoryHandler(e, categoryUsecase)
	_cartHandler.NewCartHandler(e, cartUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
