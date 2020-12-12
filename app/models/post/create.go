package post

import (
	"fmt"
	"github.com/p9c/okno/app/jdb"
	"github.com/satori/go.uuid"
	"time"
)

// Create inserts a new post to the database
func Create(j *jdb.JDB, title string, content []byte, slug string, isDraft bool) (Post, error) {
	post := Post{
		ID:        uuid.Must(uuid.NewV4(), nil).String(),
		Title:     title,
		Content:   content,
		Slug:      slug,
		IsDraft:   isDraft,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	fmt.Println(post.Title)

	err := j.Write(post.ID, post)
	return post, err
}
