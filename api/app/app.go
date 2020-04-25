package app

import (
	"api/app/db"
	"api/app/handler"
	"api/app/router"
	"api/app/store"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

type App struct {
	Echo *echo.Echo
	DB     *gorm.DB
}

func (a *App) Initialize() {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
		"utf8")

	fmt.Println("dbURI", dbURI)

	database := db.New(dbURI)
	db.AutoMigrate(database)

	r := router.New()
	api := r.Group("/api")

	cs := store.NewCompanyStore(database)
	os := store.NewOwnerStore(database)
	h := handler.NewHandler(cs, os)

	h.Register(api)

	a.Echo = r
	a.DB = database
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Echo))
}