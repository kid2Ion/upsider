package handler

import (
	"net/http"
	"upsider/usecase"

	"github.com/labstack/echo/v4"
)

type (
	invoiceHandler struct {
		usecase usecase.InvoiceUsecase
	}
	InvoiceHandler interface {
		Create() echo.HandlerFunc
		Get() echo.HandlerFunc
	}
)

func NewInvoiceHandler(usecase usecase.InvoiceUsecase) InvoiceHandler {
	return &invoiceHandler{usecase: usecase}
}

func (t *invoiceHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(usecase.InvoiceCreateReq)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		if err := t.usecase.Create(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, nil)
	}
}

func (t *invoiceHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		fromDate := c.QueryParam("from_date")
		toDate := c.QueryParam("to_date")
		res, err := t.usecase.Get(fromDate, toDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, res)
	}
}
