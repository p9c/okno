package post

import (
	"encoding/json"
	"fmt"
	"github.com/p9c/okno/app/jdb"
)

// GetAll gets all posts that match a criteria
// isDraft doesn't look at the is_draft criteria if set to -1, only finds
// non-drafts if set to 0 and only rafts if set to 1
func GetAll(j *jdb.JDB, isDraft int) (*[]Post, error) {
	var posts []Post
	ps, err := j.ReadAll()
	for _, postInterface := range ps {
		var p Post
		if err := json.Unmarshal([]byte(postInterface), &p); err != nil {
			fmt.Println("Error", err)
		}
		posts = append(posts, p)
	}

	if isDraft < 0 {
		//rows, err = conn.Query(`
		//    SELECT * FROM posts
		//    ORDER BY created_at DESC
		//`)
	} else {
		//rows, err = conn.Query(`
		//    SELECT * FROM posts
		//    WHERE is_draft = $1
		//    ORDER BY created_at DESC
		//`, isDraft)
	}
	if err != nil {
		return nil, err
	}
	//posts := []Post{}
	//for rows.Next() {
	post := Post{}
	//err := rows.Scan(
	//	&post.ID,
	//	&post.Title,
	//	&post.Content,
	//	&post.Slug,
	//	&post.IsDraft,
	//	&post.CreatedAt,
	//	&post.UpdatedAt,
	//)
	if err != nil {
		return nil, err
	}
	posts = append(posts, post)
	//}
	return &posts, nil
}
