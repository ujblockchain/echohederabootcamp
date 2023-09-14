package contract

import (
	"fmt"
	"time"

	"github.com/ujblockchain/echohederabootcamp/controller/blockchain"
	database "github.com/ujblockchain/echohederabootcamp/db"
	model "github.com/ujblockchain/echohederabootcamp/models"
)

// save deploy contract details in db
func DeployContract(productID string) model.Contract {

	//connect to database
	db, err := database.ConnectToDb()

	//if db connection fails
	if err != nil {
		panic(err)
	}

	//init contract model
	var modelContract model.Contract

	//check if a contract has been deployed,
	//we want to avoid double deployment
	dbRecord := db.First(&modelContract)

	if dbRecord.Error != nil {
		//no deployment found

		//Migrate the schema to created the contract tables
		db.AutoMigrate(&model.Contract{})

		//deploy contract
		contract, _ := blockchain.DeployContract()

		//unmarshal contract id record
		// the contract id from hedera is in pointer,
		// we need to unmarshal it to a usable form
		// this unmarshal is coming from Scan in model/contract
		cID := model.ContractID(*contract.Receipt.ContractID)

		// Create contract db record using gorm
		db.Create(&model.Contract{
			Id:            productID,
			ContractId:    &cID,
			GasUsed:       contract.CallResult.GasUsed,
			TransactionId: fmt.Sprint(contract.TransactionID),
			Timestamp:     contract.ConsensusTimestamp,
			ChargeFee:     fmt.Sprint(contract.TransactionFee),
			PayerAccount:  fmt.Sprint(contract.TransactionID.AccountID),
			Status:        fmt.Sprint(contract.Receipt.Status),
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		})

	} else {
		//call search query for the first record.
		// The first record will be the deployed contract since there will be just one record
		db.First(&modelContract)
	}

	return modelContract
}
