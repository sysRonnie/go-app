package handler

import (
	"database/sql"

	"context"
	"github.com/labstack/echo/v4"
	"go-app/model"
	"go-app/view/page"
	"net/http"
)

type UserHandler struct {
	DB *sql.DB
}

func (h UserHandler) RenderTest(c echo.Context) error {
	return render(c, page.RenderTestPage())
}

func (h UserHandler) RenderLandingPage(c echo.Context) error {
	// We pass down our user parameter
	u := model.User{
		Email: "",
	}

	// Then we render our user template, with our user data
	return render(c, page.ShowLandingPage(u))
}

// Registration handler
func (h *UserHandler) HandleRegister(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.String(http.StatusBadRequest, "Email and password cannot be empty")
	}

	// Insert new user into the database
	_, err := h.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, password)
	if err != nil {
		c.Logger().Errorf("Error inserting user into database: %v", err)
		return c.String(http.StatusInternalServerError, "Error registering user")
	}

	return c.String(http.StatusOK, "User registered successfully!")
}

// Login handler
func (h *UserHandler) HandleLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Query user from the database
	var dbPassword string
	err := h.DB.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&dbPassword)
	if err == sql.ErrNoRows {
		return c.String(http.StatusUnauthorized, "Invalid email or password")
	} else if err != nil {
		return c.String(http.StatusInternalServerError, "Error querying database")
	}

	// Verify password (in a real application, use proper hashing like bcrypt)
	if password != dbPassword {
		return c.String(http.StatusUnauthorized, "Invalid email or password")
	}

	// Set user in the context/session
	ctx := context.WithValue(c.Request().Context(), "user", email)
	c.SetRequest(c.Request().WithContext(ctx))

	return c.String(http.StatusOK, "Login successful!")
}
