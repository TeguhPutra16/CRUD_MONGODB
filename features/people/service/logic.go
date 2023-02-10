package service

import (
	"errors"
	"teguh/features/people"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PersonService struct {
	personRepository people.RepositoryEntities //data repository dri entities

}

func New(repo people.RepositoryEntities) people.ServiceEntities { //dengan kembalian user.service
	return &PersonService{
		personRepository: repo,
	}
}

// Create implements people.ServiceEntities
func (service *PersonService) Create(input people.CorePerson) (err error) {
	errCreate := service.personRepository.Create(input)
	if errCreate != nil {
		return errors.New("GAGAL MENAMBAH DATA , QUERY ERROR")
	}
	return nil
}

// GetAll implements people.ServiceEntities
func (service *PersonService) GetAll() ([]people.CorePerson, error) {
	Data, err := service.personRepository.GetAll()
	if err != nil {
		return nil, errors.New("Gagal menampilkan")
	}
	return Data, nil
}

// GetPerson implements people.ServiceEntities
func (service *PersonService) GetPerson(id primitive.ObjectID) (people.CorePerson, error) {
	Data, err := service.personRepository.GetPerson(id)
	if err != nil {
		return people.CorePerson{}, errors.New("Gagal menampilkan")
	}
	return Data, nil
}
