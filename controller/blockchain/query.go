package blockchain

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
)

// there are two function in our solidity contract contract/contract.sol
// we want to query get_record function

func GetContractRecord(newContractID hedera.ContractID, productID string) (hedera.ContractFunctionResult, error) {

	//init client and connect to it
	client, err := GetClient()

	//check if there is any error connecting to it
	if err != nil {
		return hedera.ContractFunctionResult{}, err
	}

	// Make sure to close client after running using defer
	// this will ensure it closes ony when the function ends
	// no need to keep the client connection open forever
	defer func() {
		err = client.Close()

		//check if there was an error closing the connection
		if err != nil {
			println(err.Error(), ": error closing client")
			return
		}
	}()

	// set the parameters for the contract
	// from our solidity contract contract/contract.sol, get_record requires
	// the product id parameter
	// here were are just init the parameter, we are yet to attach it to the function call
	contractFunctionParams := hedera.NewContractFunctionParameters().
		AddString(productID)

	// Call a method on a contract that exists on the contract, and attach the parameter to it
	contractQuery, err := hedera.NewContractCallQuery().
		// Set which contract id from the deployed contract
		SetContractID(newContractID).
		// Set gas to use
		SetGas(100000).
		// Set the query payment explicitly since sometimes automatic payment calculated
		// may be too low low for the transaction
		SetQueryPayment(hedera.NewHbar(1)).
		// Set the function to call on the contract
		// along with the parameter required
		SetFunction("get_record", contractFunctionParams).
		Execute(client)

	//check if the function call returned an error
	if err != nil {
		println(err.Error(), ": error executing contract call query")
		return hedera.ContractFunctionResult{}, err
	}

	//return smart contract query
	return contractQuery, err
}
