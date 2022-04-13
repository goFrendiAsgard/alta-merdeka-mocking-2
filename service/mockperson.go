package service

import "merdeka/model"

type MockPersonService struct {
	data []model.Person
}

func (ps *MockPersonService) Add(person model.Person) (model.Person, error) {
	ps.data = append(ps.data, person)
	return person, nil
}

func (ps *MockPersonService) Get() ([]model.Person, error) {
	return ps.data, nil
}

func NewMockPersonService() *MockPersonService {
	return &MockPersonService{
		data: []model.Person{},
	}
}
