package models

import (
	"example.com/tuto/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config config.Config) {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DbURL, // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,         // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	// Still debating whether to auto migrate.
	// err = database.AutoMigrate(&Branch{}, &Requirement{}, &BranchRequirement{})
	// if err != nil {
	// 	return
	// }

	DB = database
}
