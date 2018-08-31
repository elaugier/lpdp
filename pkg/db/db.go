package db

import (
	"io"
	"log"

	"github.com/elaugier/lpdp/pkg/config"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/jinzhu/gorm"

	//Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//NewInstance ...
func NewInstance(logMode bool, io io.Writer) *Instance {
	configuration, err := config.Get()
	c, err := gorm.Open("postgres", configuration.GetString("database.postgres"))
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
	}
	log.Println("database connected")
	c.LogMode(true)
	c.SetLogger(log.New(io, "", log.LstdFlags))
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

	c := i.Connection
	sql := "CREATE DATABASE IF NOT EXISTS lpdp;"
	log.Println(sql)
	c.Exec(sql)
	sql = "USE lpdp;"
	log.Println(sql)
	c.Exec("USE lpdp;")
	c.AutoMigrate(
		&models.Achievement{},
		&models.Activity{},
		&models.Alert{},
		&models.AlertAction{},
		&models.Badge{},
		&models.BadgeType{},
		&models.BadIPAddress{},
		&models.Book{},
		&models.BookPart{},
		&models.CoAuthor{},
		&models.Comment{},
		&models.Contact{},
		&models.Contest{},
		&models.ContestRound{},
		&models.CorrectionRequest{},
		&models.CorrectionRequestAction{},
		&models.IPHistory{},
		&models.JudgingPanel{},
		&models.JudgingPanelMember{},
		&models.License{},
		&models.Like{},
		&models.Message{},
		&models.Post{},
		&models.Request{},
		&models.Section{},
		&models.Tag{},
		&models.User{},
		&models.Warning{},
		&models.WarningTemplate{},
	)

	return
}

//Close ...
func (i Instance) Close() error {
	return i.Connection.Close()
}
