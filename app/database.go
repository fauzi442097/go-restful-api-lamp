package app

import (
	"fmt"
	"go-restful-api-lamp/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	connection = config.DefaultConnection
	dbConn     = config.Connections[connection]
	db         *gorm.DB
	err        error
)

func ConnectDatabase() *gorm.DB {

	switch connection {
	case "mysql":
		db, err = connectMySQL()
	case "pgsql":
		db, err = connectPostgreSQL()
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses connect ke database")

	return db
}

func connectMySQL() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConn["username"], dbConn["password"], dbConn["host"], dbConn["port"], dbConn["database"])
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func connectPostgreSQL() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbConn["host"], dbConn["username"], dbConn["password"], dbConn["database"], dbConn["port"])
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
