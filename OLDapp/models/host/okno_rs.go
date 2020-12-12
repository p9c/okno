package host

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

//func  oknoRS() *Host {
//	return &Host{
//		Host:     "okno.rs:3000",
//		Renderer: template.GetRenderer("themes/%s/views/*.html", o.Config.AppTheme),
//		Routes: func(e *echo.Echo, h handlers.Handlers) {
//            //// Public routes
//            e.GET("/", h.Homepage)
//            e.GET("/:slug", h.Page)
//            e.GET("/posts/:slug", h.Post)
//            e.POST("/api/v1/login", h.APILogin)
//		},
//	}
//}

func okno() *Host {
	//---------
	// Website
	//---------
	site := echo.New()
	site.Use(middleware.Logger())
	site.Use(middleware.Recover())
	site.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Website")
	})
	return &Host{
		Host: "okno.rs:3333",
		Echo: site,
	}
}

//
//func blog()func(db *scribble.Driver)[]*Host {
//    return func(db *scribble.Driver) []*Host{
//    //------
//    // Blog
//    //------
//    blog := echo.New()
//    blog.Use(middleware.Logger())
//    blog.Use(middleware.Recover())
//    blog.GET("/", func(c echo.Context) error {
//        return c.String(http.StatusOK, "Blog")
//    })
//    return &Host{
//        Host: "blog.okno.rs:3333",
//        Echo: blog,
//    }
//}
//}
