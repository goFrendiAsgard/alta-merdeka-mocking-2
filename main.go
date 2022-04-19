package main

import (
	"merdeka/controller"
	"merdeka/model"
	"merdeka/service"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connectionString := os.Getenv("APP_DB_CONNECTION_STRING")
	if connectionString == "" {
		connectionString = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.AutoMigrate(&model.Person{})

	ps := service.NewDBPersonService(db)
	pc := controller.NewPersonController(ps)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World 🌏")
	})
	e.GET("/persons", pc.Get)  // mendapatkan persons
	e.POST("/persons", pc.Add) // menambah person
	e.Logger.Fatal(e.Start(":" + port))
}
