package db

import (
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/elaugier/lpdp/pkg/config"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/jinzhu/gorm"

	//Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Inst ...
var Inst Instance

//GetInstance ...
func GetInstance() *gorm.DB {
	if Inst != (Instance{}) {
		return Inst.c
	}
	Inst = *NewInstance(true, gin.DefaultWriter)
	Inst.c.DB().SetMaxIdleConns(10)
	Inst.c.DB().SetMaxOpenConns(100)
	Inst.c.DB().SetConnMaxLifetime(time.Hour)
	return Inst.c
}

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
	i := Instance{
		c: c,
	}
	i.DatabaseInitialization()
	return &i
}

//Instance ...
type Instance struct {
	//c ...
	c *gorm.DB
}

//DatabaseInitialization ...
func (i Instance) DatabaseInitialization() {

	c := i.c
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
		&models.Notification{},
		&models.Payment{},
		&models.Post{},
		&models.Rating{},
		&models.Read{},
		&models.Request{},
		&models.Section{},
		&models.Tag{},
		&models.User{},
		&models.Vote{},
		&models.Voting{},
		&models.Warning{},
		&models.WarningTemplate{},
	)

	return
}

//Close ...
func (i Instance) Close() error {
	return i.c.Close()
}
