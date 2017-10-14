package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	cfg "github.com/NgulikinApp/luffy/config"
	"github.com/labstack/echo"

	"github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"

	mware "github.com/NgulikinApp/luffy/middleware"
	userHandler "github.com/NgulikinApp/luffy/user/delivery/http"
	userRepo "github.com/NgulikinApp/luffy/user/repository/mysql"
	userUcase "github.com/NgulikinApp/luffy/user/usecase"

	categoryHandler "github.com/NgulikinApp/luffy/category/delivery/http"
	categoryRepo "github.com/NgulikinApp/luffy/category/repository/mysql"
	categoryUcase "github.com/NgulikinApp/luffy/category/usecase"

	productHandler "github.com/NgulikinApp/luffy/product/delivery/http"
	productRepo "github.com/NgulikinApp/luffy/product/repository/mongodb"
	productUcase "github.com/NgulikinApp/luffy/product/usecase"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if config.GetBool(`debug`) {
		logrus.Warn(`Luffy is running in debug mode`)
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func main() {
	dbHost := config.GetString(`database.host`)
	dbPort := config.GetString(`database.port`)
	dbUser := config.GetString(`database.user`)
	dbPass := config.GetString(`database.pass`)
	dbName := config.GetString(`database.name`)

	dsn := dbUser + `:` + dbPass + `@tcp(` + dbHost + `:` + dbPort + `)/` + dbName + `?parseTime=1&loc=Asia%2FJakarta`
	logrus.Info("connecting to database")
	db, err := sql.Open(`mysql`, dsn)
	if err != nil {
		logrus.Error(fmt.Sprintf("database connection failed. Err: %v", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	duration, err := time.ParseDuration(config.GetString(`mongodb.timeout`))
	if err != nil {
		logrus.Error(fmt.Sprintf("Failed to parse duration. Err: %v", err.Error()))
	}
	info := &mgo.DialInfo{
		Addrs:    []string{config.GetString(`mongodb.address`)},
		Timeout:  duration * time.Second,
		Database: config.GetString(`mongodb.database`),
		Username: config.GetString(`mongodb.username`),
		Password: config.GetString(`mongodb.password`),
		Source:   config.GetString(`mongodb.source`),
	}

	mongoUrl := config.GetString(`mongodb.url`)
	mongoName := config.GetString(`mongodb.name`)
	logrus.Debug(`Using MongoDB: `, mongoUrl+mongoName)
	mongoSession, err := mgo.DialWithInfo(info)
	if err != nil {
		logrus.Error(fmt.Sprintf("MongoDB connection failed. Err: %v", err.Error()))
		os.Exit(1)
	}

	e := echo.New()
	e.Use(mware.SetCORS())

	e.GET(`/ping`, func(c echo.Context) error {
		return c.String(http.StatusOK, `kapan soft launch?`)
	})

	userRepository := new(userRepo.MySQLRepository)
	userRepository.Conn = db
	userUsecase := userUcase.NewUsecase(userRepository)
	userHandler.Init(e, userUsecase)

	categoryRepository := new(categoryRepo.MySQLRepository)
	categoryRepository.Conn = db
	categoryUsecase := categoryUcase.NewUsecase(categoryRepository)
	categoryHandler.Init(e, categoryUsecase)

	productRepository := productRepo.NewRepository(mongoSession, mongoName)
	productUsecase := productUcase.NewUsecase(productRepository)
	productHandler.Init(e, productUsecase)

	address := config.GetString(`server.address`)
	logrus.Infof(`Luffy server running at address : %v`, address)

	e.Start(config.GetString(`server.address`))
}
