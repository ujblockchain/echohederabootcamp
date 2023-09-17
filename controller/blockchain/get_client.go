package blockchain

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/ujblockchain/echohederabootcamp/config"
)

// set environment
var env = config.EnVar

// func to set and get hedera client
func GetClient() (*hedera.Client, error) {

	//init client and error var
	// needed to connect to hedera network
	var client *hedera.Client
	var err error

	// Retrieving network type from environment variable HEDERA_NETWORK
	// we are using the test network
	client, err = hedera.ClientForName(env.GetString("HEDERA_NETWORK"))

	//check if connection was established
	if err != nil {
		println(err.Error(), ": error creating client")

		// we return an empty client and an error
		return &hedera.Client{}, err
	}

	// Retrieving account ID from environment variable in development.yaml
	operatorAccountID, err := hedera.AccountIDFromString(env.GetString("ACCOUNT_ID"))

	//check if there was an error getting the account in
	if err != nil {
		println(err.Error(), ": error converting string to AccountID")

		// we return an empty client and an error
		return &hedera.Client{}, err
	}

	// Retrieving private key from environment variable in development.yaml
	operatorKey, err := hedera.PrivateKeyFromString(env.GetString("DER_ENCODED_PRIVATE_KEY"))
	if err != nil {
		println(err.Error(), ": error converting string to PrivateKey")

		// we return an empty client and an error
		return &hedera.Client{}, err
	}

	// connect to the client using your account ID and private key
	client.SetOperator(operatorAccountID, operatorKey)

	return client, nil

}
