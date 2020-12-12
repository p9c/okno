package app

//func (o *OKNO) tester() *host.Host {
//	return &host.Host{
//		Host:     "test.okno.rs:1323",
//		Renderer: template.GetRenderer("themes/%s/views/*.html", o.Config.AppTheme),
//		Routes: func(e *echo.Echo,hs handlers.Handlers) {
//            //// Auth-only v1 API routes
//            e.GET("/", func(c echo.Context) error {
//                return c.String(http.StatusOK, "Tester")
//            })
//            connStr := fmt.Sprintf(":%d", o.Config.AppPort)
//            //Start the server
//            e.Logger.Fatal(e.Start(connStr))
//		},
//	}
//}
//
//
