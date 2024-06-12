package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"harmoniq/harmoniq-api-v2/pkg/ehttp"
	"harmoniq/harmoniq-api-v2/setup"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	loadConfig()
}

func loadConfig() {
	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	db := setupDatabase()
	defer closeDatabaseConnection(db)

	e := setupServer()

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	repos := setup.NewRepositories(db)
	useCases := setup.NewUseCases(repos, time.Duration(timeoutContext)*time.Second)
	setup.SetupHandlers(e, useCases)

	log.Fatal(e.Start(viper.GetString("server.address")))
}

func setupDatabase() *gorm.DB {
	connectionString := buildConnectionString()
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

func buildConnectionString() string {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	return fmt.Sprintf("%s?%s", connection, val.Encode())
}

func closeDatabaseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sqlDB from gorm DB:", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatal("Failed to close database connection:", err)
	}
}

func setupServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header"},
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())
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
	return e
}
