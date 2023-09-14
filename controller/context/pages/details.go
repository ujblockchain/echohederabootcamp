package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echohederabootcamp/db/product"
)

func DetailsContext(c echo.Context) error {
	//get product id from query parameter
	getProductId := c.Param("productId")

	// you can also print if you wish to see it
	// fmt.Println(getProductId)

	//get db record
	dbRecord, _ := product.GetProductDetails(getProductId)

	//render template but pass the found db record: "dbRecord"
	// note tha "dbRecord" is of map[string]interface{}{}
	// which is exactly what this context requires
	return c.Render(http.StatusOK, "detail.html", dbRecord)
}
