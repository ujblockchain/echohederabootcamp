package db

import (
	"github.com/ujblockchain/echohederabootcamp/config"
	//driver that will allow to work with postgres db
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// init env variable
var env = config.EnVar

func ConnectToDb() (*gorm.DB, error) {
	//connect to db using the DNS configuration in development.yaml
	db, err := gorm.Open(postgres.Open(env.GetString("DNS")), &gorm.Config{})

	//check if we were able to connect to database
	if err != nil {
		panic("fail to connect to database")
	}

	//return the connected database
	return db, err
}
