package main

import (
	"context"
	"fmt"
	"go-app/db"
	"go-app/handler"
	"github.com/labstack/echo/v4"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize the database and handle any errors
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.Close() // Ensure the DB is closed when the application exits

	app := echo.New()

    app.Static("/", "public")

	// Import the UserHandler Struct from the handler
	userHandler := handler.UserHandler{DB: db} // Passing DB to UserHandler
	app.Use(withUser)

	// Define routes
	app.GET("/user", userHandler.HandleUserShow)
	app.POST("/login", userHandler.HandleLogin)
	app.POST("/register", userHandler.HandleRegister)

	// Start the server
	if err := app.Start(":3000"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Working..")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "a@gg.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

