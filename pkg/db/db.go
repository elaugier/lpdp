package db

import (
	"log"

	"github.com/elaugier/lpdp/pkg/config"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/jinzhu/gorm"
	//Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//NewInstance ...
func NewInstance() *Instance {
	configuration, err := config.Get()
	c, err := gorm.Open("postgres", configuration.GetString("database.postgres"))
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
	}
	log.Println("database connected")
	return &Instance{
		Connection: c,
	}
}

//Instance ...
type Instance struct {
	Connection *gorm.DB
}

//DatabaseInitialization ...
func (i Instance) DatabaseInitialization() {

	i.Connection.Exec("CREATE DATABASE IF NOT EXISTS lpdp;")
	i.Connection.Exec("USE lpdp;")

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
