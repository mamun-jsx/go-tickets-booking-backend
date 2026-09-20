package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/mamun-jsx/go-tickets-booking-backend.git/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
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
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)
	e.POST("/users", userHandler.CreateUser)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
