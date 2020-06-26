package infrastructure

import (
	"fmt"
	"prototype2/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

// SetupModels : initializing mysql database
func SetupModels() *gorm.DB {
	get := utils.GetEnvWithKey
	USER := get("DB_USER")
	PASS := get("DB_PASS")
	HOST := get("DB_HOST")
	PORT := get("DB_PORT")
	DBNAME := get("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	fmt.Println(URL)
	db, err := gorm.Open("mysql", URL)

	if err != nil {
		panic(err.Error())
	}

	return db
}

/**
* Only for the purpose of integratino testing
*/
func SetupModelsForControllerTest() *gorm.DB {
	USER := "root"
	PASS := "password"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "prototype"

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)

	if err != nil {
		panic(err.Error())
	}

	return db
}
