package page

import (
	"time"

	"github.com/p9c/okno/OLDapp/models/post"
)

type Page struct {
	ID           string
	PostID       string
	Index        int
	Slug         string
	InNavigation bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Post         post.Post
}
