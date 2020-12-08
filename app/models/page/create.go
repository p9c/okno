package page

import (
	"github.com/1lann/cete"
	postModel "github.com/p9c/okno/app/models/post"
	"github.com/satori/go.uuid"
	"time"
)

// Create inserts a new page to the DB
func Create(db *cete.DB, postID string, index int, slug string, inNavigation bool) (*Page, error) {
	page := Page{
		ID:           uuid.Must(uuid.NewV4(), nil).String(),
		PostID:       postID,
		Index:        index,
		Slug:         slug,
		InNavigation: inNavigation,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Post:         postModel.Post{},
	}

	//err := conn.QueryRow(`
	//    INSERT INTO
	//    pages (post_id, index, slug, in_navigation, created_at, updated_at)
	//    VALUES
	//    ($1, $2, $3, $4, NOW(), NOW())
	//    RETURNING *
	//    `,
	//	postID,
	//	index,
	//	slug,
	//	inNavigation,
	//).Scan(
	//	&page.ID,
	//	&page.PostID,
	//	&page.Index,
	//	&page.Slug,
	//	&page.InNavigation,
	//	&page.CreatedAt,
	//	&page.UpdatedAt,
	//)
	//if err != nil {
	//	return nil, err
	//}
	db.Table("pages").Set(postID, page)
	return &page, nil
}