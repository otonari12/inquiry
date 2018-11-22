package inquiry

import (
	"net/http"

	"github.com/labstack/echo"
)

func Get() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "Get Hello World")
	}
}

func Post() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "Post Hello World")
	}
}
