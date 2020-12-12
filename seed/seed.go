package seed

import (
	"github.com/p9c/okno/app/jdb"
	"github.com/p9c/okno/app/models/page"
	"github.com/p9c/okno/app/models/post"
)

func Seed(j *jdb.JDB) error {
	posts, err := seedPosts(j)
	if err != nil {
		return err
	}

	_, err = page.Create(j, (*posts)[0].ID, 1, "about-my-vacation", true)
	if err != nil {
		return err
	}
	return nil
}

func seedPosts(j *jdb.JDB) (*[]post.Post, error) {
	posts := []post.Post{}
	for _, p := range getPosts() {
		post, err := post.Create(
			j,
			p.Title,
			p.Content,
			p.Slug,
			p.IsDraft,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}
