package post

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

// GetByID returns a single post with the ID specified
func GetByID(db *scribble.Driver, id string) (*Post, error) {
	post := Post{}
	//err := conn.QueryRow(
	//	`SELECT * FROM posts WHERE id = $1`,
	//	id,
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
	//	return nil, err
	//}

	//db.Table("posts").Get(id, &post)

	return &post, nil
}
