package delivery

import (
	"net/http"
	"teguh/features/people"
	"teguh/helper"

	"github.com/labstack/echo/v4"
)

type PersonDelivery struct {
	PersonService people.ServiceEntities
}

func New(Service people.ServiceEntities, e *echo.Echo) {
	handler := &PersonDelivery{
		PersonService: Service,
	}

	e.POST("/person", handler.Create)
	// e.GET("/users", handler.GetAll, middlewares.JWTMiddleware())
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
