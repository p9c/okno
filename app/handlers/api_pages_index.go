package handlers

import (
	"time"

	"github.com/labstack/echo/v4"
)

// APIPagesIndex shows all pages
func (hs Handlers) APIPagesIndex(c echo.Context) error {
	type ResponseItemPost struct {
		ID             string    `json:"id"`
		Title          string    `json:"title"`
		ContentPreview string    `json:"content_preview"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
	type ResponseItem struct {
		ID           string           `json:"id"`
		PostID       string           `json:"post_id"`
		Index        int              `json:"index"`
		Slug         string           `json:"slug"`
		InNavigation bool             `json:"in_navigation"`
		CreatedAt    time.Time        `json:"created_at"`
		UpdatedAt    time.Time        `json:"updated_at"`
		Post         ResponseItemPost `json:"post"`
	}
	//pages, err := pageModel.GetAll(hs.DB, true)
	//if err != nil {
	//	return err
	//}
	//res := []ResponseItem{}
	//for _, page := range *pages {
	//	page.Post.ContentPreview = page.Post.GetContentPreview(80)
	//	item := ResponseItem{
	//		ID:           page.ID,
	//		PostID:       page.PostID,
	//		Index:        page.Index,
	//		Slug:         page.Slug,
	//		InNavigation: page.InNavigation,
	//		CreatedAt:    page.CreatedAt,
	//		UpdatedAt:    page.UpdatedAt,
	//		Post: ResponseItemPost{
	//			ID:             page.Post.ID,
	//			Title:          page.Post.Title,
	//			ContentPreview: page.Post.ContentPreview,
	//			CreatedAt:      page.Post.CreatedAt,
	//			UpdatedAt:      page.Post.UpdatedAt,
	//		},
	//	}
	//	res = append(res, item)
	//}
	//return c.JSON(http.StatusOK, res)
	return nil
}
