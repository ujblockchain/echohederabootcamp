package constant

import (
	"github.com/labstack/echo/v4"
)

func LoadStatic(app *echo.Echo) {
	//load favicon
	app.File("/favicon.ico", "repository/assets/images/favicon.ico")

	//get path to static file like css, js fonts, images etc
	app.Static("static", "repository/assets")
}
