package app

import (
	"github.com/labstack/echo/v4"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/p9c/okno/OLDapp/jdb"
	"github.com/p9c/okno/OLDapp/models/host"
	"log"

	"github.com/p9c/okno/OLDapp/config"
)

type OKNO struct {
	Config config.Config
	Db     *scribble.Driver
	Hosts  map[string]*echo.Echo
}

func NewOKNO() *OKNO {
	config, err := config.GetConfiguration()
	if err != nil {
		log.Fatal(err)
	}
	//Get the database connection
	db, err := jdb.InitDB("./cete_data")
	if err != nil {
		log.Fatal(err)
	}

	okno := &OKNO{
		Config: config,
		Db:     db,
		Hosts:  host.Server(config, db),
	}

	return okno
}
