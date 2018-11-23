package api

import (
	"inquiry-api/inquiry"
	"net/http"

	"github.com/labstack/echo"
)

// Get Metdod
func Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, inquiry.List())
	}
}

// Post Method
func Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		item := new(inquiry.InquiryItem)
		if err := c.Bind(item); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		item.Save()

		return c.NoContent(http.StatusOK)
	}
}
