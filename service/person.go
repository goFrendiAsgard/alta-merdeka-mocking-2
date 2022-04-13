package service

import (
	"merdeka/model"
)

type PersonService interface {
	Add(person model.Person) (model.Person, error)
	Get() ([]model.Person, error)
}
