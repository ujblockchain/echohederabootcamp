package product

import (
	"fmt"
	"time"

	database "github.com/ujblockchain/echohederabootcamp/db"
	"github.com/ujblockchain/echohederabootcamp/models"
)

func CreateRecord(productID string, cornColor string, quantity int64, timestamp time.Time) error {
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

	//declare what record to save in the products database table
	db.Create(&models.Products{
		ProductID: productID,
		Color:     cornColor,
		Quantity:  quantity,
		Timestamp: timestamp,
	})

	return nil
}
