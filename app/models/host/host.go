package host

import (
	"github.com/labstack/echo/v4"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/p9c/okno/OLDapp/config"
)

type Host struct {
	Host string
	Echo *echo.Echo
	//Renderer *template.Template
	//Routes   func(e *echo.Echo, h handlers.Handlers)
}

//func getHosts(h []*Host) map[string]*echo.Echo {
//	// Hosts
//	hosts := map[string]*echo.Echo{}
//
//	//-----
//	// API
//	//-----
//
//	hosts["api.localhost:1323"] = &Host{api}
//
//	api.GET("/", func(c echo.Context) error {
//		return c.String(http.StatusOK, "API")
//	})
//
//}

//func (h *Host) host(hs handlers.Handlers) {
//	// Init Echo
//	h.Echo = echo.New()
//	// Set our template renderer as the renderer for Echo
//	h.Echo.Renderer = h.Renderer
//	h.Echo.Use(middleware.Logger())
//	h.Echo.Use(middleware.Recover())
//	h.Routes(h.Echo, hs)
//}

//func Server(hosts []*Host, hand handlers.Handlers) map[string]*echo.Echo {
//	// Hosts
//	hostEcho := make(map[string]*echo.Echo)
//	for _, h := range hosts {
//		h.host(hand)
//		hostEcho[h.Host] = h.Echo
//	}
//	return hostEcho
//}

func Server(c config.Config, db *scribble.Driver) map[string]*echo.Echo {
	return map[string]*echo.Echo{
		okno().Host:     okno().Echo,
		api(c, db).Host: api(c, db).Echo,
		admin(c).Host:   admin(c).Echo,
	}
}
