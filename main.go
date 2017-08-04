package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	cfg "github.com/NgulikinApp/luffy/config"
	"github.com/labstack/echo"

	"github.com/Sirupsen/logrus"

	userHandler "github.com/NgulikinApp/luffy/user/delivery/http"
	uRepo "github.com/NgulikinApp/luffy/user/repository/mysql"
	uUcase "github.com/NgulikinApp/luffy/user/usecase"
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

	e := echo.New()

	e.GET(`/ping`, func(c echo.Context) error {
		return c.String(http.StatusOK, `pong`)
	})

	userRepository := new(uRepo.UserRepository)
	userRepository.Conn = db
	userUsecase := uUcase.NewUserUsecase(userRepository)
	userHandler.Init(e, userUsecase)

	address := config.GetString(`server.address`)
	logrus.Infof(`Luffy server running at address : %v`, address)

	e.Start(config.GetString(`server.address`))
}
