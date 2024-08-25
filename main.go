package main

import (
	"upsider/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// TODO: 認証やロギング、CORS設定などは適宜middlewareとして追加する

	router.InitInvoiceRouter(e)
	// http://localhost:8080
	e.Start(":8080")
}
