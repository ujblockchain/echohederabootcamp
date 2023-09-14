package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echohederabootcamp/controller/context/pages"
)

func IndexRouter(app *echo.Echo) {
	app.GET("/", pages.IndexContext)
}
