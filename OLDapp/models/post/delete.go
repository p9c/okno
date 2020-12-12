package post

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

// Delete removes a post from the database and returns it
func Delete(db *scribble.Driver, id string) (Post, error) {
	post := Post{}
	//err := conn.QueryRow(`DELETE FROM posts WHERE id = $1 RETURNING *`, id).Scan(
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
	//db.Table("posts").Delete(id, 0)

	return post, nil
}
