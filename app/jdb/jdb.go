package jdb

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

type JDB struct {
	db   *scribble.Driver
	path string
}

func NewJDB(db *scribble.Driver, path string) *JDB {
	return &JDB{
		db:   db,
		path: path,
	}
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// DB is the central database for jorm
//var DB, _ = scribble.New(cfg.Dir, nil)

// // ReadData appends 'data' path prefix for a database read
func (j *JDB) Read(id string) (data interface{}, err error) {
	err = j.db.Read(j.path, id, &data)
	return
}

// Read all items from the database, unmarshaling the response.
func (j *JDB) ReadAll() (data []interface{}, err error) {
	records, err := j.db.ReadAll(j.path)
	if err != nil {
		fmt.Println("Error", err)
	}
	for _, d := range records {
		data = append(data, d)
	}
	return
}

// Write appends post path prefix for a database write
func (j *JDB) Write(id string, data interface{}) (err error) {
	//err = j.db.Write(j.path, id, data)
	if err := j.db.Write(j.path, id, data); err != nil {
		fmt.Println("Error", err)
	}
	return
}
