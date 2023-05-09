package db

import (
	"fmt"
	"log"
	"shiftsync/pkg/config"
	"shiftsync/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatbase(config config.Config) *gorm.DB {
	connstr := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.Db.DbHost, config.Db.DbUser, config.Db.DbName, config.Db.DbPort, config.Db.DbPaswword)
	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
		return nil
	}

	if err := db.AutoMigrate(
		domain.Employee{},
		domain.Attendance{},
		domain.Form{},
		domain.Admin{},
		domain.Leave{},
	); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected succesfully....!")

	return db
}
