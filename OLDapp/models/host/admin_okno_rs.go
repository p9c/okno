package host

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/p9c/okno/OLDapp/config"
	"net/http"
)

func admin(c config.Config) *Host {
	//---------
	// Website
	//---------
	admin := echo.New()
	admin.Use(middleware.Logger())
	admin.Use(middleware.Recover())
	//admin.Static("/admin", "public")
	admin.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin")
	})
	admin.Static("/admin", "public")
	//admin.Static(
	//    fmt.Sprintf("/%s", c.AppTheme),
	//    fmt.Sprintf("./themes/%s/public", c.AppTheme),
	//)
	return &Host{
		Host: "admin.okno.rs:3333",
		Echo: admin,
	}
}
