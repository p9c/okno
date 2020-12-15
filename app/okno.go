package app

import (
	"fmt"
	"github.com/gorilla/handlers"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/p9c/okno/app/config"
	"github.com/p9c/okno/app/models/app"
	"html/template"
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
		//Handler: o.Handlers(),
		//Handler: interceptHandler(o.Handlers(), defaultErrorHandler),
		Handler: handlers.CORS()(handlers.CompressHandler(interceptHandler(o.Handlers(), defaultErrorHandler))),
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

type interceptResponseWriter struct {
	http.ResponseWriter
	errH func(http.ResponseWriter, int)
}

func (w *interceptResponseWriter) WriteHeader(status int) {
	if status >= http.StatusBadRequest {
		w.errH(w.ResponseWriter, status)
		w.errH = nil
	} else {
		w.ResponseWriter.WriteHeader(status)
	}
}

type ErrorHandler func(http.ResponseWriter, int)

func (w *interceptResponseWriter) Write(p []byte) (n int, err error) {
	if w.errH == nil {
		return len(p), nil
	}
	return w.ResponseWriter.Write(p)
}

func defaultErrorHandler(w http.ResponseWriter, status int) {
	//    const tpl = `
	//<!DOCTYPE html>
	//<html>
	//	<head>
	//		<meta charset="UTF-8">
	//	</head>
	//	<body>
	//		<div>{{ . }}</div>
	//
	//	</body>
	//</html>`
	t := template.Must(template.ParseFiles("js/errors/error.html"))
	//check := func(err error) {
	//    if err != nil {
	//        log.Fatal(err)
	//    }
	//}

	//t, err := template.New("webpage").Parse(tpl)
	//check(err)

	//err := template.GetRenderer("./js/errors/error.html", "").Render(w, "error", map[string]interface{}{"status": status})

	//tmpl := template.Must(template.ParseFiles("js/errors/error.html"))
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t.Execute(w, map[string]interface{}{"status": status})
	//})
}

func interceptHandler(next http.Handler, errH ErrorHandler) http.Handler {
	if errH == nil {
		errH = defaultErrorHandler
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(&interceptResponseWriter{w, errH}, r)
	})
}
