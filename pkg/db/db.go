package db

import (
	"fmt"
	"log"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/jinzhu/gorm"
	//Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Instance ...
type Instance struct {
	Driver     string
	Hostname   string
	Port       string
	DbName     string
	Username   string
	Password   string
	Connection *gorm.DB
}

//DatabaseInitialization ...
func (i Instance) DatabaseInitialization() {

	if (!i.Connection.HasTable(&models.Achievement{})) {
		i.Connection.CreateTable(&models.Achievement{})
	}

	if (!i.Connection.HasTable(&models.Activity{})) {
		i.Connection.CreateTable(&models.Activity{})
	}

	return
}

//Connect ...
func (i Instance) Connect() error {
	var err error
	i.Connection, err = gorm.Open(i.Driver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		i.Hostname, i.Port, i.Username, i.DbName, i.Password))
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
		return err
	}
	log.Println("database connected")
	return nil
}

//Close ...
func (i Instance) Close() error {
	return i.Connection.Close()
}
