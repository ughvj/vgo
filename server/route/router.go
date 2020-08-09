package route

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{ "http://localhost:8080" },
		AllowHeaders: []string{ echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept },
	}))

	r := e.Group("/hello")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", hello)

	return e;
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}