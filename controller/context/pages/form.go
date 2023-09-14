package pages

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echohederabootcamp/controller/blockchain"
	"github.com/ujblockchain/echohederabootcamp/db/contract"
	"github.com/ujblockchain/echohederabootcamp/db/product"
)

func FormContext(c echo.Context) error {
	//get user product id input
	productID := c.FormValue("product_id")
	//get user product quantity input
	quantity := c.FormValue("quantity")
	//get user product color input
	cornColor := c.FormValue("corn_color")

	//convert quantity to number from string
	// to int64, same used the in the smart contract
	numQuantity, _ := strconv.ParseInt(quantity, 10, 64)

	// read contract
	// deploy contract if it has not been deployed before
	record := contract.DeployContract(productID)
	hederaContractID := hedera.ContractID(*record.ContractId)

	//save record in hedera blockchain
	blockchainRecord, err := blockchain.SetContractRecord(hederaContractID, productID, cornColor, numQuantity)

	//check is there was error with the save
	if err != nil {
		panic(err) // panic stops the application
	}

	//get contract id of transaction
	contractId := blockchainRecord.Receipt.ContractID // returns a pointer

	//get string format of contraction id
	contractIdStr := fmt.Sprint(blockchainRecord.Receipt.ContractID)

	//get the gas used for the transaction
	gasUsed := blockchainRecord.CallResult.GasUsed

	//get transaction id for the transaction
	transactionId := fmt.Sprint(blockchainRecord.TransactionID)

	//transaction timestamp
	timestamp := blockchainRecord.ConsensusTimestamp
	
	//transaction fee
	chargeFee := fmt.Sprint(blockchainRecord.TransactionFee)

	// fee payer account
	payerAccount := fmt.Sprint(blockchainRecord.TransactionID.AccountID)

	//status of transaction
	status := fmt.Sprint(blockchainRecord.Receipt.Status)

	//save product record in db
	err = product.CreateRecord(productID, cornColor, numQuantity, timestamp, contractId, contractIdStr, gasUsed, transactionId, chargeFee, payerAccount, status)

	//check if there is error saving in the database
	if err != nil {
		panic(err)
	}

	//set path
	path := fmt.Sprintf("/%v", productID)

	//direct to details page
	return c.Redirect(http.StatusMovedPermanently, path)
}
