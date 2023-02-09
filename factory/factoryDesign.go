package factory

import (
	PersonDelivery "teguh/features/people/delivery"
	PersonRepo "teguh/features/people/repository"
	PersonService "teguh/features/people/service"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitFactory(e *echo.Echo, client *mongo.Client) {
	PersonRepofaktory := PersonRepo.New(client)
	PersonServicefaktory := PersonService.New(PersonRepofaktory)
	PersonDelivery.New(PersonServicefaktory, e)

}
