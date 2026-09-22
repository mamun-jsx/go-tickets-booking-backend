package user

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/mamun-jsx/go-tickets-booking-backend.git/internal/httpresponse"
	"github.com/mamun-jsx/go-tickets-booking-backend.git/internal/user/dto"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

// create user

func (h *handler) CreateUser(c *echo.Context) error {
	var req dto.CreateRequest // user input

	// * input from user
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{Code: http.StatusBadRequest, Message: "Invalid Input", Details: err.Error()})
	}
	// validate the input
	if err := c.Validate(&req); err != nil {

		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation Failed",
			Details: err.Error(),
		})
	}

	res, err := h.service.CreateUser(req)

	if errors.Is(err, ErrorAlreadyExist) {
		return c.JSON(http.StatusConflict, httpresponse.Error{
			Code:    http.StatusConflict,
			Message: "Try another email",
			Details: ErrorAlreadyExist.Error(),
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Details: err.Error(),
		})
	}
	// success response
	return c.JSON(http.StatusCreated, res)
}