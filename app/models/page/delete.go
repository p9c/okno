package page

import (
	"github.com/1lann/cete"
)

// Delete removes a page from the DB
func (p Page) Delete(db *cete.DB) error {
	//_, err := conn.Exec("DELETE FROM pages WHERE id = $1", p.ID)
	//if err != nil {
	//	return err
	//}
	db.Table("pages").Delete(p.ID, 0)
	return nil
}
