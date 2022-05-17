package test

import (
	"fmt"
	"go-restful-api-lamp/repositories"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	port     = "5432"
	database = "lamp"
	username = "postgres"
	password = "''"
)

func setupDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, username, password, database, port)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db
}

func TestGetAll(t *testing.T) {

	db := setupDB()

	tx := db.Begin()
	customerRepository := repositories.NewCustomerRepository()
	customers, _ := customerRepository.GetAll(tx)

	for i, customer := range customers {
		fmt.Println("No.",i+1, " ID: ", customer.ID, " ", customer.FullName)
	}
}
