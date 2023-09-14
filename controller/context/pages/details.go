package pages

import (
	"fmt"
	"net/http"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echohederabootcamp/controller/blockchain"
	"github.com/ujblockchain/echohederabootcamp/db"
	"github.com/ujblockchain/echohederabootcamp/models"
)

func DetailsContext(c echo.Context) error {
	//get product id from query parameter
	getProductId := c.Param("productId")

	//get db record using the product id
	// first connect to db
	db, err := db.ConnectToDb()

	//return error if connect fails
	if err != nil {
		fmt.Println("Failed to connect to database")
	}

	//init the product model
	var product models.Products

	//search db for product id that equals to getProductId
	result := db.Where("product_id = ?", getProductId).First(&product)

	//if record is found, return it
	if result.Error == nil {
		// product now has the found record
		//get the contract id from the database
		// and unmarshal it using scan(automatically) in models/contract
		hederaContractID := hedera.ContractID(*product.ContractId)
		//pass the unmarshal contract id and product id to query the smart contract in hedera
		// and return the query result
		contractQuery, err := blockchain.GetContractRecord(hederaContractID, getProductId)

		//check if there is an error
		if err != nil {
			panic(err) // if error, panic stops the application
		}

		//pass query result as context
		return c.Render(http.StatusOK, "detail.html", map[string]interface{}{
			"product_id":      product.ProductID,
			"contract_id":     product.ContractIdString,
			"gas_used":        contractQuery.GasUsed,
			"transaction_fee": product.ChargeFee,
			"payer_account":   product.PayerAccount,
		})
	}

	return err
}
