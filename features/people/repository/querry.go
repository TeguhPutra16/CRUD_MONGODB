package repository

import (
	"teguh/features/people"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	client *mongo.Client
}

func New(client *mongo.Client) people.RepositoryEntities { // user.repository mengimplementasikan interface repository yang ada di entities
	return &userRepository{
		client: client,
	}

}

// Create implements people.RepositoryEntities
func (*userRepository) Create(input people.CorePerson) (err error) {
	panic("unimplemented")
}
