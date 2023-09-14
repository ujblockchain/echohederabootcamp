package product

import (
	"fmt"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
	database "github.com/ujblockchain/echohederabootcamp/db"
	"github.com/ujblockchain/echohederabootcamp/models"
)

func CreateRecord(productID string, cornColor string, quantity int64, timestamp time.Time, contractId *hedera.ContractID, contractIdString string, gasUsed uint64, transactionId string, chargeFee string, payerAccount string, status string) error {
	// try to connect to database using DNS config in
	// environment variables
	db, err := database.ConnectToDb()

	//if database is not found, return error
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	// Migrate the database schema
	// this will create the table products if it does not exit
	db.AutoMigrate(&models.Products{})

	//this will make use of the unmarshal and unmarshal
	cID := models.ContractID(*contractId)

	//declare what record to save in the products database table
	db.Create(&models.Products{
		ProductID:        productID,
		Color:            cornColor,
		Quantity:         quantity,
		Timestamp:        timestamp,
		ContractId:       &cID,
		ContractIdString: contractIdString,
		GasUsed:          gasUsed,
		TransactionId:    transactionId,
		ChargeFee:        chargeFee,
		PayerAccount:     payerAccount,
		Status:           status,
		CreatedAt:        time.Now(),
	})

	return nil
}
