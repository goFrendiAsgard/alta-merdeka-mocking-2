package controller

import (
	"fmt"
	"merdeka/model"
	"merdeka/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PersonController struct {
	ps service.PersonService
}

func (pc PersonController) Get(c echo.Context) error {
	persons, err := pc.ps.Get()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "cannot get persons"})
	}
	return c.JSON(http.StatusOK, persons)
}

func (pc PersonController) Add(c echo.Context) error {
	var newPerson model.Person
	c.Bind(&newPerson)
	person, err := pc.ps.Add(newPerson)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "cannot get persons"})
	}
	return c.JSON(http.StatusOK, person)
}

func NewPersonController(ps service.PersonService) PersonController {
	return PersonController{
		ps: ps,
	}
}
