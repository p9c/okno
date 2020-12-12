package template

import (
	"io"

	"github.com/labstack/echo/v4"
)

// Render handles the actual rendering of the templates
func (t *Template) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
