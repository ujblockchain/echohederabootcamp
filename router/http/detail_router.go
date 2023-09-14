package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echohederabootcamp/controller/context/pages"
)

func DetailsRouter(app *echo.Echo) {
	app.GET("/:productId", pages.DetailsContext)
}
