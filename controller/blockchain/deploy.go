package blockchain

import (
	"encoding/json"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

// we need to extract the byte code from the  compiled solidity contract at contract/contract.json
// we use struct to first target data, bytecode then object base on how the contract/contract.json file struct
// Thus bytecode struct will be the obe with the actual bytecode object
// the object contains the byte code contract

// contract data struct
type ContractData struct {
	Data Data `json:"data"`
}

type Data struct {
	Bytecode Bytecode `json:"bytecode"`
}

// contract bytecode struct
type Bytecode struct {
	Object string `json:"object"`
}

//  Store the Smart Contract Bytecode on Hedera

func DeployContract() (*hedera.TransactionRecord, error) {
	// we first need to connect to the hedera client using your private
	// key and account ID in development.yaml

	//init client and check for error error
	client, err := GetClient()
	if err != nil {
		return &hedera.TransactionRecord{}, err
	}

	// Read in the compiled contract from the contract.json file in contract/contract.json
	rawSmartContract, err := os.ReadFile("contract/contract.json")

	//check if there was error reading the file
	if err != nil {
		println(err.Error(), ": error reading contract.json")
		return &hedera.TransactionRecord{}, err
	}

	// Initialize contracts with is of Bytecode struct down used to extract the bytecode object
	var contract ContractData

	// Parse the bytecode object in contract.json into contract since it if ContractData
	// which allows the copying of the bytecode object
	err = json.Unmarshal([]byte(rawSmartContract), &contract)

	//check if there was error copying it
	if err != nil {
		println(err.Error(), ": error unmarshaling")
		return &hedera.TransactionRecord{}, err
	}

	//Create the transaction
	// create file storing the bytecode and contract
	// set the amount of gas needed for the transaction
	// then part the bytecode object to it
	contractCreate := hedera.NewContractCreateFlow().
		SetGas(100000).
		SetBytecode([]byte(contract.Data.Bytecode.Object))

	//Sign the transaction with the hedera client key created using your
	//private key and submit to a Hedera network
	txResponse, err := contractCreate.Execute(client)

	//check if there was any error signing the transaction
	if err != nil {
		return &hedera.TransactionRecord{}, err
	}

	//get the transaction record of the contract deployed
	getRecord, err := txResponse.GetRecord(client)
	if err != nil {
		return &hedera.TransactionRecord{}, err
	}

	//use pointer for the record
	record := &getRecord

	//return the record and nil since there will be no error at this point
	return record, nil
}
