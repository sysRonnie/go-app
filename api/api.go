package api

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserAPI struct {
	DB *sql.DB
}

func (h UserAPI) Test(c echo.Context) error {
	return c.String(http.StatusOK, "Hello DOOOOOOD! WHATS UP!!!!! KEEP GOING")
}

// Registration handler

func (h *UserAPI) HandleRegister(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.String(http.StatusBadRequest, "<p>Email and password cannot be empty</p>")
	}

	_, err := h.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, password)
	if err != nil {
		c.Logger().Errorf("Error inserting user into database: %v", err)
		return c.String(http.StatusInternalServerError, "<p>Error registering user</p>")
	}

	return c.String(http.StatusOK, "<p>User registered successfully!</p>")
}

// Login handler
func (h *UserAPI) HandleLogin(c echo.Context) error {
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
