package config

import (
	"MyGram/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	// db configuration
	Username string
	Password string
	Port     string
	Address  string
	Database string

	// db connection
	DB *gorm.DB
}

type GormDb struct {
	*Postgres
}

var (
	GORM *GormDb
)

func InitPostgres() error {
	GORM = new(GormDb)

	GORM.Postgres = &Postgres{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	// connect to database
	err := GORM.Postgres.OpenConnection()
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) OpenConnection() error {
	// init dsn
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Address, p.Port, p.Username, p.Password, p.Database)

	dbConnection, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return err
	}

	p.DB = dbConnection

	err = p.DB.Debug().AutoMigrate(model.User{}, model.Comment{}, model.Photo{}, model.SocialMedia{})

	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to database")

	return nil
}
