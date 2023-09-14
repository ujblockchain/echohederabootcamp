package blockchain

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
)

func SetContractRecord(newContractID hedera.ContractID, productID string, color string, quantity int64) (*hedera.TransactionRecord, error) {

	//init client and error var
	client, err := GetClient()
	if err != nil {
		return &hedera.TransactionRecord{}, err
	}

	// Make sure to close client after running using defer
	// this will ensure it closes ony when the function ends
	// no need to keep the client connection open forever
	defer func() {
		err = client.Close()
		if err != nil {
			println(err.Error(), ": error closing client")
			return
		}
	}()

	// // set the parameters for the contract
	// from our solidity contract contract/contract.sol, set_record requires
	// the product id parameter
	// here were are just init the parameter, we are yet to attach it to the function call
	contractFunctionParams := hedera.NewContractFunctionParameters().
		AddString(productID).
		AddString(color).
		AddInt64(quantity)

	// add record
	contractExecuteID, err := hedera.NewContractExecuteTransaction().
		// Set which contract id from the deployed contract
		SetContractID(newContractID).
		// Set the gas to execute the contract call
		SetGas(2000000).
		// Set the function to call and the parameters to send
		// in this case we're calling function "set_record" with a
		// parameters for product id, corn colors, and quantity"
		SetFunction("set_record", contractFunctionParams).
		Execute(client)

	//check if the function call returned an error
	if err != nil {
		println(err.Error(), ": error executing contract")
		return &hedera.TransactionRecord{}, err
	}

	// Retrieve the record to make sure the execute transaction ran
	contractExecuteRecord, err := contractExecuteID.GetRecord(client)

	//check if there was an error
	if err != nil {
		println(err.Error(), ": error retrieving contract execution record")
		return &hedera.TransactionRecord{}, err
	}

	// get transaction record
	record := &contractExecuteRecord

	//
	return record, nil
}
