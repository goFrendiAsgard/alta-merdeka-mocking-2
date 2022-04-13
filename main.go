package main

import (
	"merdeka/controller"
	"merdeka/model"
	"merdeka/service"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.AutoMigrate(&model.Person{})

	ps := service.NewDBPersonService(db)
	pc := controller.NewPersonController(ps)
	e.GET("/persons", pc.Get)  // mendapatkan persons
	e.POST("/persons", pc.Add) // menambah person
	e.Logger.Fatal(e.Start(":8000"))
}
