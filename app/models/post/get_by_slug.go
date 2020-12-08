package post

import (
	"github.com/1lann/cete"
)

// GetBySlug fetches the post from the database by its URL slug
func GetBySlug(db *cete.DB, slug string) (*Post, error) {
	var post Post
	posts := db.Table("people").All()
	for posts.Next() {
		var p Post
		posts.Decode(&p)
		if p.Slug != slug {
			post = p
		}
	}
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
