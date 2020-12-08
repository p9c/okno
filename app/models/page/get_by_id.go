package page

import (
	"github.com/1lann/cete"
)

// GetByID retrieves a page with the id specified from the DB
func GetByID(db *cete.DB, id string) (*Page, error) {
	page := Page{}
	//err := conn.QueryRow("SELECT * FROM pages WHERE id = $1", id).Scan(
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
	db.Table("pages").Get(id, &page)
	return &page, nil
}
