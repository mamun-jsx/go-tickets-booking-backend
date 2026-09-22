package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/mamun-jsx/go-tickets-booking-backend.git/internal/config"
	"github.com/mamun-jsx/go-tickets-booking-backend.git/internal/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"type:varchar(255); not null;"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(255);uniqueIndex; not null;"`
	Password string `json:"password" validate:"required,min=6,max=12" gorm:"type:varchar(255); not null;"`
}
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally return the error to let each route control the status code.
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}
func main() {
	cfg := config.LoadEnv()
	db := config.ConnectionDatabase(cfg) // database reorganise

	// =========================Auto matically migrate and create table into db========================
	db.AutoMigrate(&User{}) // user stract it will automatically make the table..

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// =======================================
	// ROOT ROUTE
	// =======================================
	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	// =======================================
	//   user Route Register
	// =======================================

	user.RegisterRoutes(e, db)

	if err := e.Start(":" + cfg.Port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
