package delivery

import (
	"net/http"
	"teguh/features/people"
	"teguh/helper"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PersonDelivery struct {
	PersonService people.ServiceEntities
}

func New(Service people.ServiceEntities, e *echo.Echo) {
	handler := &PersonDelivery{
		PersonService: Service,
	}

	e.POST("/person", handler.Create)
	e.GET("/people", handler.GetAll)
	e.GET("/person/:id", handler.GetPerson)
	// e.PUT("/users/:id", handler.Update, middlewares.JWTMiddleware())
	// e.DELETE("/users/:id", handler.DeleteById, middlewares.JWTMiddleware())
	// e.GET("/users/:id", handler.GetById, middlewares.JWTMiddleware())

}

func (delivery *PersonDelivery) Create(c echo.Context) error {
	InputPerson := PersonRequest{} //penangkapan data user reques dari entities user
	errbind := c.Bind(&InputPerson)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errbind.Error()))
	}

	dataCore := PersonRequestToPersonCore(InputPerson) //data mapping yang diminta create
	errResultCore := delivery.PersonService.Create(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errResultCore.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success create person"))
}

func (delivery *PersonDelivery) GetAll(c echo.Context) error {
	result, err := delivery.PersonService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadGateway, helper.FailedResponse("erorr read data"+err.Error()))

	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get People", result))
}

func (delivery *PersonDelivery) GetPerson(c echo.Context) error {
	id, errReq := primitive.ObjectIDFromHex(c.Param("id"))
	if errReq != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errReq.Error()))
	}
	result, err := delivery.PersonService.GetPerson(id)
	if err != nil {
		return c.JSON(http.StatusBadGateway, helper.FailedResponse("erorr read data"+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get Persin", result))

}
