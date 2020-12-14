package app

import (
	"fmt"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/p9c/okno/app/config"
	"github.com/p9c/okno/app/models/app"
	"net/http"
	"time"
)

const (
	// HTTPMethodOverrideHeader is a commonly used
	// http header to override a request method.
	HTTPMethodOverrideHeader = "X-HTTP-Method-Override"
	// HTTPMethodOverrideFormKey is a commonly used
	// HTML form key to override a request method.
	HTTPMethodOverrideFormKey = "_method"
)

func NewOKNO() *app.OKNO {
	conf, err := config.GetConfiguration()
	// a new scribble driver, providing the directory where it will be writing to,
	// and a qualified logger if desired
	db, err := scribble.New("./"+conf.DBName, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	o := &app.OKNO{
		Database: db,
	}

	srv := &http.Server{
		Handler: o.Handlers(),
		//Handler: handlers.CORS()(handlers.CompressHandler(o.Handlers())),
		// Handler: cacheHandler(handlers.CORS()(handlers.CompressHandler(r))),
		// Handler: handlers.CompressHandler(r),
		Addr: ":" + conf.AppPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	o.Server = srv
	return o
}
