package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echohederabootcamp/router/http"
)

func LoadAllRouters(app *echo.Echo) {
	//index router
	http.IndexRouter(app)

	//form router
	http.FormRouter(app)

	//details details
	http.DetailsRouter(app)
}
