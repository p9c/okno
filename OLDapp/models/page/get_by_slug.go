package page

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

// GetBySlug gets the page with the slug specified from the DB
func GetBySlug(db *scribble.Driver, slug string) (*Page, error) {
	page := Page{}
	//db.Table("pages").Index("Age").One(10, &page)
	//err := jdb.QueryRow("SELECT * FROM pages WHERE slug = $1", slug).Scan(
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
	return &page, nil
}
