package people

import "go.mongodb.org/mongo-driver/bson/primitive"

type CorePerson struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

type ServiceEntities interface {
	Create(input CorePerson) (err error)
	GetAll() ([]CorePerson, error)
	GetPerson(id primitive.ObjectID) (CorePerson, error)
}

type RepositoryEntities interface {
	Create(input CorePerson) (err error)
	GetAll() ([]CorePerson, error)
	GetPerson(id primitive.ObjectID) (CorePerson, error)
}
