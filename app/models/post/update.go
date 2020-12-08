package post

import (
	"github.com/1lann/cete"
	"time"
)

// Update the DB record for a post
func Update(db *cete.DB, id string, title string, content string, slug string, isDraft bool) (Post, error) {
	post := Post{
		ID:        id,
		Title:     title,
		Content:   content,
		Slug:      slug,
		IsDraft:   isDraft,
		UpdatedAt: time.Now(),
	}
	db.Table("posts").Update(id, post)
	//err := conn.QueryRow(`
	//    UPDATE posts
	//    SET title = $2, content = $3, slug = $4, is_draft = $5, updated_at = NOW()
	//    WHERE id = $1
	//    RETURNING *
	//    `,
	//	id,
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
	return post, nil
}
