package post

import (
	"github.com/1lann/cete"
)

// GetAllByIds gets only the posts with IDs specified
func GetAllByIds(db *cete.DB, ids []string) (*[]Post, error) {
	// Create a SQL param string from the IDs provided
	//idsString := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
	//param := "{" + idsString + "}"
	//rows, err := conn.Query("SELECT * FROM posts WHERE id = ANY($1::int[])", param)
	//if err != nil {
	//	return nil, err
	//}
	posts := []Post{}

	for _, id := range ids {
		post := Post{}
		db.Table("posts").Get(id, &post)
		posts = append(posts, post)
	}
	//for rows.Next() {
	//	post := Post{}
	//	err := rows.Scan(
	//		&post.ID,
	//		&post.Title,
	//		&post.Content,
	//		&post.Slug,
	//		&post.IsDraft,
	//		&post.CreatedAt,
	//		&post.UpdatedAt,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}
	//	posts = append(posts, post)
	//}
	return &posts, nil
}
