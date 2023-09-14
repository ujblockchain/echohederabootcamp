package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"gorm.io/gorm"
)

// init hedera contract id
type ContractID hedera.ContractID

// contract model
type Contract struct {
	gorm.Model
	Id            string
	ContractId    *ContractID
	GasUsed       uint64
	TransactionId string
	Timestamp     time.Time
	ChargeFee     string
	PayerAccount  string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// method to unmarshal the contract id data
// the contract id from hedera is in pointer,
// we need to unmarshal it to a usable form
func (c *ContractID) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal ContractID value:", value))
	}

	return json.Unmarshal(bytes, c)
}

// method to marshal the contract id data
// this is used to convert an contract id back into a suitable hedera form of memory pointer
func (c ContractID) Value() (driver.Value, error) {
	return json.Marshal(c)
}
