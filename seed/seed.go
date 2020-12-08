package seed

import (
	"github.com/1lann/cete"

	pageModel "github.com/p9c/okno/app/models/page"
)

func Seed(db *cete.DB) error {
	posts, err := seedPosts(db)
	if err != nil {
		return err
	}
	_, err = pageModel.Create(db, (*posts)[0].ID, 1, "about-my-vacation", true)
	if err != nil {
		return err
	}
	return nil
}
