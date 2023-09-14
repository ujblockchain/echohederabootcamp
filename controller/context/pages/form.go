package pages

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	database "github.com/ujblockchain/echohederabootcamp/db/product"
)

func FormContext(c echo.Context) error {
	//get user product id input
	productID := c.FormValue("product_id")
	//get user product quantity input
	quantity := c.FormValue("quantity")
	//get user product color input
	cornColor := c.FormValue("corn_color")

	//convert quantity to number from string
	numQuantity, _ := strconv.ParseInt(quantity, 10, 64)

	//get current time
	currentTime := time.Now()

	//create databse record
	database.CreateRecord(productID, cornColor, numQuantity, currentTime)

	//set path
	path := fmt.Sprintf("/%v", productID)

	//direct to details page
	return c.Redirect(http.StatusMovedPermanently, path)
}
