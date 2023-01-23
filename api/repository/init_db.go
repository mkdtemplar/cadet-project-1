package repository

import (
	"cadet-project/configurations"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (u *PG) InitDb() {
	var err error
	configurations.InitDbConfig("configurations")

	DbUrl := configurations.ConfigDB.ConnectionString()

	u.DB, err = gorm.Open(postgres.Open(DbUrl), &gorm.Config{})

	if err != nil {
		fmt.Printf("Cannot connect to %s database %s", configurations.ConfigDB.DBDriver, configurations.ConfigDB.DBName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database: %s\n", configurations.ConfigDB.DBDriver, configurations.ConfigDB.DBName)
	}

}
