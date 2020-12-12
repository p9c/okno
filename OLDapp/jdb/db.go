package jdb

import (
	"github.com/nanobox-io/golang-scribble"
)

func InitDB(dir string) (*scribble.Driver, error) {
	// a new scribble driver, providing the directory where it will be writing to,
	// and a qualified logger if desired
	db, err := scribble.New(dir, nil)
	//if err != nil {
	//    fmt.Println("Error", err)
	//}
	return db, err
}
