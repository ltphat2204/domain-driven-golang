package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func (c *DBConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
}

func GetDBConfig() (*DBConfig, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, fmt.Errorf("DB_HOST is not set")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return nil, fmt.Errorf("DB_USER is not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is not set")
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		return nil, fmt.Errorf("DB_NAME is not set")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return nil, fmt.Errorf("DB_PORT is not set")
	}

	return &DBConfig{
		Host:     host,
		User:     user,
		Password: password,
		Name:     name,
		Port:     port,
	}, nil
}

func GetDb() (*gorm.DB, error) {
	dbConfig, err := GetDBConfig()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dbConfig.ConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}