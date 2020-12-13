package app

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/p9c/okno/app/hosts"
	"github.com/p9c/okno/app/jdb"
	"github.com/p9c/okno/seed"
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

type OKNO struct {
	Database *scribble.Driver
	Hosts    []*hosts.Host
	Server   *http.Server
}

func NewOKNO() *OKNO {
	// a new scribble driver, providing the directory where it will be writing to,
	// and a qualified logger if desired
	db, err := scribble.New("./DATABASE", nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	o := &OKNO{
		Database: db,
	}

	var r = mux.NewRouter()
	for _, host := range hosts.GetHosts(o.Database) {
		host.Routes(r)

		seed.Seed(jdb.NewJDB(o.Database, host.Slug))
	}
	srv := &http.Server{
		Handler: handlers.CORS()(handlers.CompressHandler(r)),
		// Handler: cacheHandler(handlers.CORS()(handlers.CompressHandler(r))),
		// Handler: handlers.CompressHandler(r),
		Addr: ":4433",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	o.Server = srv
	return o
}

// func cacheHandler(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var age time.Duration
// 		ext := filepath.Ext(r.URL.String())

// 		// Timings are based on github.com/h5bp/server-configs-nginx

// 		switch ext {
// 		case ".rss", ".atom":
// 			age = time.Hour / time.Second
// 		case ".css", ".js":
// 			age = (time.Hour * 24 * 365) / time.Second
// 		case ".jpg", ".jpeg", ".gif", ".png", ".ico", ".cur", ".gz", ".svg", ".svgz", ".mp4", ".ogg", ".ogv", ".webm", ".htc":
// 			age = (time.Hour * 24 * 30) / time.Second
// 		default:
// 			age = 1
// 		}

// 		if age > 0 {
// 			w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d, public, must-revalidate, proxy-revalidate", age))
// 			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
// 		}

// 		h.ServeHTTP(w, r)
// 	})
// }
