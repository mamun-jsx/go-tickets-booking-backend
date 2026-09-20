package main

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
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
	dsn := "postgresql://neondb_owner:npg_5rXJsiavH1QP@ep-rapid-hat-b4hepzjn-pooler.c-6.us-east-2.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic("Failed to conncet databae")
	} else {
		println("===================================")
		println("______Database connected______")
		println("===================================")
	}
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
	// create a user
	// =======================================
	e.POST("/users", func(c *echo.Context) error {
		newUser := new(User) // new always make an empty strat and return the pointer
		// ? binding the user data
		if err := c.Bind(newUser); err != nil {
			return err
		}

		// validating user data
		if err := c.Validate(newUser); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{"message": err.Error()})
		}

		// * save to database
		result := db.Create(newUser) // save to database
		if result.Error != nil {      // check the error
			if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
				return c.JSON(http.StatusConflict, map[string]any{"message": "Email already exists"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{"error": result.Error.Error()})
		}

		return c.JSON(http.StatusOK, newUser)
	})
	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}