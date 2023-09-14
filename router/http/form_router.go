package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echohederabootcamp/controller/context/pages"
)

func FormRouter(app *echo.Echo) {
	app.POST("/record", pages.FormContext)
}
