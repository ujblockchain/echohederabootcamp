package models

import (
	"time"

	"gorm.io/gorm"
)

// declare the database table and the record
// the table should always contain
type Products struct {
	gorm.Model
	ProductID        string
	Color            string
	Quantity         int64
	Timestamp        time.Time
	ContractId       *ContractID
	ContractIdString string
	GasUsed          uint64
	TransactionId    string
	ChargeFee        string
	PayerAccount     string
	Status           string
	ReceiptPDF       string
	CreatedAt        time.Time
}
