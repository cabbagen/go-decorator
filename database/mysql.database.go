package database

import (
	"cts-go/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var databaseHandler *gorm.DB

func Connect() {
	connectString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", config.DatabaseConfig["username"], config.DatabaseConfig["password"], config.DatabaseConfig["dbname"])

	dbf, error := gorm.Open("mysql", connectString)

	if error != nil {
		log.Fatal(error)
	}
	databaseHandler = dbf
}

func Destroy() {
	if databaseHandler == nil {
		return
	}
	databaseHandler.Close()
}

func GetDatabaseHandleInstance() *gorm.DB {
	if databaseHandler == nil {
		fmt.Sprintln("you should connect the database before invoke `GetDatabaseHandleInstance` method")
	}
	return databaseHandler
}

