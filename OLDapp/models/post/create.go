package post

import (
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/satori/go.uuid"
	"time"
)

// Create inserts a new post to the database
func Create(db *scribble.Driver, title string, content string, slug string, isDraft bool) (Post, error) {
	post := Post{
		ID:        uuid.Must(uuid.NewV4(), nil).String(),
		Title:     title,
		Content:   content,
		Slug:      slug,
		IsDraft:   isDraft,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	//err := conn.QueryRow(`
	//    INSERT INTO posts
	//        (title, content, slug, is_draft, created_at, updated_at)
	//    VALUES ($1, $2, $3, $4, NOW(), NOW())
	//    RETURNING *
	//    `,
	//	title,
	//	content,
	//	slug,
	//	isDraft,
	//).Scan(
	//	&post.ID,
	//	&post.Title,
	//	&post.Content,
	//	&post.Slug,
	//	&post.IsDraft,
	//	&post.CreatedAt,
	//	&post.UpdatedAt,
	//)
	//if err != nil {
	//	return Post{}, err
	//}
	//db.Table("posts").Set(post.ID, post)

	return post, nil
}
