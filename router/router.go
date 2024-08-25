package router

import (
	"upsider/di"

	"github.com/labstack/echo/v4"
)

func InitInvoiceRouter(e *echo.Echo) {
	handler := di.InjectInvoiceHandler()
	e.POST("/api/invoices", handler.Create())
}
