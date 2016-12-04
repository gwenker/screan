package server

import (
	"net/http"
	"strconv"
	
	"github.com/labstack/echo"
)

type Banana struct {
	One string `json:"one,omitempty"`
}

// Start Start server
func Start(port int) {
	e := echo.New()
	e.Debug = true
	e.POST("/", func(c echo.Context) error {
		var data Banana
		c.Bind(&data)

		return c.JSON(http.StatusOK, data)
	})
	e.Logger.Fatal(e.Start(":"+strconv.Itoa(port)))
}
