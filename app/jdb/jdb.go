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
	col  string
}

func NewJDB(db *scribble.Driver, path string) *JDB {
	return &JDB{
		db:   db,
		path: path,
	}
}

// // ReadData appends 'data' path prefix for a database read
func (j *JDB) Collection(c string) (err error) {
	j.col = c
	return
}

// // ReadData appends 'data' path prefix for a database read
func (j *JDB) Read(id string) (data interface{}, err error) {
	err = j.db.Read(j.path+"/"+j.col, id, &data)
	return
}

// Read all items from the database, unmarshaling the response.
func (j *JDB) ReadAll() (data []string, err error) {
	data, err = j.db.ReadAll(j.path + "/" + j.col)
	if err != nil {
		fmt.Println("Error", err)
	}
	return
}

// Write appends post path prefix for a database write
func (j *JDB) Write(id string, data interface{}) (err error) {
	if err := j.db.Write(j.path+"/"+j.col, id, data); err != nil {
		fmt.Println("Error", err)
	}
	return
}

// Delete  data from the database
func (j *JDB) Delete(id string) (data interface{}, err error) {
	if err = j.db.Delete(j.path+"/"+j.col, id); err != nil {
		fmt.Println("Error", err)
	}
	return
}
