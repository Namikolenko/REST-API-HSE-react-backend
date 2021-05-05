package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

var db *gorm.DB
var tokenPaswd tokenPassword

type dbConfig struct {
	Config 		config `json:"postgres"`
}

type config struct {
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	DbName   	string `json:"dbName"`
	DbHost   	string `json:"dbHost"`
}

type tokenPassword struct {
	Info		string `json:"token_password"`
}

func (cfg config) getDsn() string {

	dsn := ""
	dsn = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.DbHost, cfg.Username, cfg.DbName, cfg.Password)

	return dsn
}

func init() {

	configJson, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	dbBytes, _ := ioutil.ReadAll(configJson)

	var dbcfg dbConfig

	err = json.Unmarshal(dbBytes, &dbcfg)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(dbBytes, &tokenPaswd)
	if err != nil {
		panic(err)
	}
	configJson.Close()


	conn, err := gorm.Open(postgres.Open(dbcfg.Config.getDsn()), &gorm.Config{})
	if err != nil {
		fmt.Println("Wrong configuration settings")
		panic(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{}, &Task{})
}

func GetDB() *gorm.DB {
	return db
}

func GetTokenPassword() string {
	return tokenPaswd.Info
}
