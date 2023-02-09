package repository

import (
	"context"
	"teguh/features/people"
	"time"

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
func (repo *userRepository) Create(input people.CorePerson) (err error) {
	PersonModel := FromCoreToModel(input)
	collection := repo.client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = collection.InsertOne(ctx, PersonModel)

	if err != nil {
		return err
	}
	return nil
}
