package page

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

// GetAll returns all pages from the database
func GetAll(db *scribble.Driver, withPosts bool) (*[]Page, error) {
	//rows, err := jdb.Query("SELECT * FROM pages ORDER BY index DESC")
	//if err != nil {
	//	return nil, err
	//}
	//rows := db.Table("pages").All()

	pages := []Page{}
	//postIds := []string{}
	//for rows.Next() {
	//	page := Page{}
	//err := rows.Scan(
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
	//rows.Decode(&page)
	//pages = append(pages, page)
	//postIds = append(postIds, page.PostID)
	//}
	//if withPosts {
	//posts, err := post.GetAllByIds(jdb, postIds)
	//if err != nil {
	//	return nil, err
	//}
	//pagesWithPosts := []Page{}
	//for _, page := range pages {
	//	pageWithPost := page
	//	for _, post := range *posts {
	//		if page.PostID == post.ID {
	//			pageWithPost.Post = post
	//		}
	//	}
	//	pagesWithPosts = append(pagesWithPosts, pageWithPost)
	//}
	//pages = pagesWithPosts
	//}
	return &pages, nil
}
