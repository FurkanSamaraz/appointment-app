package postgres

import (
	"log"
	"meeting_app/configs"
	api_structure "meeting_app/internal/app/structures"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	*gorm.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	dbConnectionString := configs.EnvGetURI("DATABASE_CONNECTION")

	db, err := gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresDB{DB: db}, nil
}

func MigrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&api_structure.Service{}, &api_structure.Appointment{}, &api_structure.Customer{}, &api_structure.CustomerServiceHistory{})
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
