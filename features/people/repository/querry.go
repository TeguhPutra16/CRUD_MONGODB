package repository

import (
	"context"
	"teguh/features/people"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// GetAll implements people.RepositoryEntities
func (repo *userRepository) GetAll() ([]people.CorePerson, error) {
	var people []Person
	collection := repo.client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	dataCore := ListModelToCore(people)
	return dataCore, nil

}
