package service

import (
	"merdeka/model"

	"gorm.io/gorm"
)

type DBPersonService struct {
	db *gorm.DB
}

func (ps DBPersonService) Add(person model.Person) (model.Person, error) {
	tx := ps.db.Save(&person)
	err := tx.Error
	return person, err
}

func (ps DBPersonService) Get() ([]model.Person, error) {
	persons := []model.Person{}
	tx := ps.db.Find(&persons)
	err := tx.Error
	return persons, err
}

func NewDBPersonService(db *gorm.DB) DBPersonService {
	return DBPersonService{
		db: db,
	}
}
