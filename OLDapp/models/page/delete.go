package page

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

// Delete removes a page from the DB
func (p Page) Delete(db *scribble.Driver) error {
	//_, err := conn.Exec("DELETE FROM pages WHERE id = $1", p.ID)
	//if err != nil {
	//	return err
	//}
	//db.Table("pages").Delete(p.ID, 0)
	return nil
}
