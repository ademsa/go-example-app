package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		hostname, _ := os.Hostname()
		return c.JSON(http.StatusOK, map[string]interface{}{
			"lang":     "go",
			"hostname": hostname,
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}
