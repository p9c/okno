package handlers

// Post renders a page with a post
func (hs Handlers) Post() error {
	hs.Collection("posts")
	//post, err := postModel.GetBySlug(hs.DB, c.Param("slug"))
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		return c.String(http.StatusNotFound, "Not found")
	//	}
	//	return err
	//}
	//pages, err := pageModel.GetAll(hs.DB, true)
	//if err != nil {
	//	return err
	//}
	//return c.Render(
	//	http.StatusOK,
	//	"post.html",
	//	map[string]interface{}{"post": post, "pages": pages},
	//)
	return nil
}
