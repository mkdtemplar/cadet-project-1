package repository

import (
	"cadet-project/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (u *PG) InitDb() {
	var err error
	config.InitDbConfig("configurations")

	DbUrl := config.ConfigDB.ConnectionString()

	u.DB, err = gorm.Open(postgres.Open(DbUrl), &gorm.Config{})

	if err != nil {
		fmt.Printf("Cannot connect to %s database %s", config.ConfigDB.DBDriver, config.ConfigDB.DBName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database: %s\n", config.ConfigDB.DBDriver, config.ConfigDB.DBName)
	}

}
