package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	pageModel "github.com/p9c/okno/OLDapp/models/page"
	postModel "github.com/p9c/okno/OLDapp/models/post"
)

// Page renders a navigable page that's not just a post
func (hs Handlers) Page(c echo.Context) error {
	page, err := pageModel.GetBySlug(hs.DB, c.Param("slug"))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "Not found")
		}
		return err
	}
	post, err := postModel.GetByID(hs.DB, page.PostID)
	if err != nil {
		return err
	}
	pages, err := pageModel.GetAll(hs.DB, true)
	if err != nil {
		return err
	}
	return c.Render(
		http.StatusOK,
		"page.html",
		map[string]interface{}{"post": post, "pages": pages},
	)
}
