package handlers

import (
	pageModel "github.com/p9c/okno/app/models/page"
	postModel "github.com/p9c/okno/app/models/post"
	"net/http"

	"github.com/labstack/echo"
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

	//fmt.Println("posts", posts)
	//fmt.Println("pages", pages)
	return c.Render(
		http.StatusOK,
		"homepage.html",
		map[string]interface{}{"posts": posts, "pages": pages},
	)
}
