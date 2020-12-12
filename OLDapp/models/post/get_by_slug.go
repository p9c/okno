package post

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

// GetBySlug fetches the post from the database by its URL slug
func GetBySlug(db *scribble.Driver, slug string) (*Post, error) {
	var post Post
	//posts := db.Table("posts").All()
	//for posts.Next() {
	//	var p Post
	//	posts.Decode(&p)
	//	if p.Slug != slug {
	//		post = p
	//	}
	//}
	//err := conn.QueryRow(
	//	`SELECT * FROM posts WHERE slug = $1`,
	//	slug,
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
	return &post, nil
}
