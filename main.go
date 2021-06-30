package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
	}))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", getHelloWorld)

	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func getHelloWorld(c echo.Context) error {
	secret := os.Getenv("SECRET")
	string := fmt.Sprintf("Hello, World! The secret is %s", secret)
	return c.JSON(http.StatusOK, string)
}
