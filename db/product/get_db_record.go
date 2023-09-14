package product

import (
	"fmt"

	database "github.com/ujblockchain/echohederabootcamp/db"
	"github.com/ujblockchain/echohederabootcamp/models"
)

func GetProductDetails(productID string) (map[string]interface{}, error) {
	//connect to db
	db, err := database.ConnectToDb()

	//return error if connect fails
	if err != nil {
		fmt.Println("Failed to connect to database")
		return map[string]interface{}{}, err
	}

	//init the product model
	var product models.Products

	//search db for product id that equals the one passed
	result := db.Where("product_id = ?", productID).First(&product)

	//if record is found, return it
	if result.Error == nil {
		return map[string]interface{}{
			"product_id": product.ProductID,
			"color":      product.Color,
			"quantity":   product.Quantity,
			"timestamp":  product.Timestamp,
		}, nil
	}

	//else return error that the id does not exit
	return map[string]interface{}{}, err
}
