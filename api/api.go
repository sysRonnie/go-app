package api

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserAPI struct {
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func (h UserAPI) Test(c echo.Context) error {
	return c.String(http.StatusOK, "Hello DOOOOOOD! WHATS UP!!!!! KEEP GOING")
}
