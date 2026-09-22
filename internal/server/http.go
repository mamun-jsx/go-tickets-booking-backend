package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/mamun-jsx/go-tickets-booking-backend.git/internal/config"
	"github.com/mamun-jsx/go-tickets-booking-backend.git/internal/user"
	"gorm.io/gorm"
	"net/http"
)

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

// server start
func Start(db *gorm.DB, cfg *config.Config) {
	db.AutoMigrate(&user.User{}) // user stract it will automatically make the table..

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
	user.RegisterRoutes(e, db)

	if err := e.Start(":" + cfg.Port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
