package db

import (
	"fmt"
	"log"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/jinzhu/gorm"
	//Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//NewInstance ...
func NewInstance(
	Driver string,
	Hostname string,
	Port string,
	DbName string,
	Username string,
	Password string) *Instance {
	c, err := gorm.Open(Driver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		Hostname, Port, Username, DbName, Password))
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
	}
	log.Println("database connected")
	return &Instance{
		Driver:     Driver,
		Hostname:   Hostname,
		Port:       Port,
		DbName:     DbName,
		Username:   Username,
		Password:   Password,
		Connection: c,
	}
}

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

//Close ...
func (i Instance) Close() error {
	return i.Connection.Close()
}
