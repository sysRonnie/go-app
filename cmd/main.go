package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"go-app/api"
	"go-app/db"
	"go-app/handler"
	"log"
)

func main() {
	// First we should initialize our database connection. I like the idea of local storage. I want to store as much of it as I can on the client side so it functions offline.
	// I think HTMX is going to do this super well. The idea is to use a noSQL solution like pocketbase for local storage and then SQL for the mainframe.
	// This stack is fucking awesome.
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.Close() // Ensure the DB is closed when the application exits

	app := echo.New()

	app.Static("/", "public")

	// Import the UserHandler Struct from the handler
	userHandler := handler.UserHandler{DB: db} // Passing DB to UserHandler
	api := api.UserAPI{}
	app.Use(withUser)

	// Define Template Routes
	app.GET("/user", userHandler.RenderLandingPage)
	app.GET("/hello", api.Test)
	// Define API Routes
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
		// This is where we should implement the user-interface
		ctx := context.WithValue(c.Request().Context(), "user", "a@gg.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
