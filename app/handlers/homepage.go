package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	pageModel "github.com/p9c/okno/app/models/page"
	postModel "github.com/p9c/okno/app/models/post"
)

// Homepage renders the index page
func (hs Handlers) Homepage(c echo.Context) error {
	posts, err := postModel.GetAll(hs.DB, 0)
	if err != nil {
		return err
	}
	pages, err := pageModel.GetAll(hs.DB, true)
	if err != nil {
		return err
	}
	return c.Render(
		http.StatusOK,
		"homepage.html",
		map[string]interface{}{"posts": posts, "pages": pages},
	)
}
