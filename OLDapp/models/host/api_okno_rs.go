package host

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/p9c/okno/OLDapp/config"
	"github.com/p9c/okno/OLDapp/handlers"
	"github.com/p9c/okno/OLDapp/template"
	"net/http"
)

func api(c config.Config, db *scribble.Driver) *Host {
	//-----
	// API
	//-----
	// Init Echo
	api := echo.New()
	// Set our template renderer as the renderer for Echo
	api.Renderer = template.GetRenderer("themes/%s/views/*.html", c.AppTheme)
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	api.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "API")
	})
	//// Init the handlers object
	hs := handlers.Handlers{
		DB:     db,
		Config: c,
	}

	//// Auth-only v1 API routes
	v1auth := api.Group("/api/v1")
	v1auth.Use(middleware.JWT(c.AppSecret))
	v1auth.GET("/posts", hs.APIPostsIndex)
	v1auth.GET("/posts/titles", hs.APIPostsTitles)
	v1auth.GET("/posts/:id", hs.APIPostsGet)
	v1auth.POST("/posts", hs.APIPostsCreate)
	v1auth.PUT("/posts/:id", hs.APIPostsUpdate)
	v1auth.DELETE("/posts/:id", hs.APIPostsDelete)
	v1auth.GET("/pages", hs.APIPagesIndex)
	v1auth.POST("/pages", hs.APIPagesCreate)
	v1auth.PUT("/pages/:id", hs.APIPagesUpdate)
	v1auth.DELETE("/pages/:id", hs.APIPagesDelete)
	//connStr := fmt.Sprintf(":%d", c.AppPort)
	////Start the server
	//api.Logger.Fatal(api.Start(connStr))
	//
	return &Host{
		Host: "api.okno.rs:3333",
		Echo: api,
	}
}
