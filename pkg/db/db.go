package db

import (
	"io"
	"log"
	"time"

	"github.com/elaugier/lpdp/pkg/logs"

	"github.com/gin-gonic/gin"

	"github.com/elaugier/lpdp/pkg/config"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/jinzhu/gorm"

	//Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Inst ...
var Inst Instance

var logger *log.Logger

//GetInstance ...
func GetInstance() *gorm.DB {
	if Inst != (Instance{}) {
		return Inst.c
	}
	Inst = *NewInstance(true, gin.DefaultWriter)
	conn := Inst.c
	conn.DB().SetMaxIdleConns(10)
	conn.DB().SetMaxOpenConns(100)
	conn.DB().SetConnMaxLifetime(time.Hour)
	return conn
}

//NewInstance ...
func NewInstance(logMode bool, io io.Writer) *Instance {
	logger = logs.GetInstance()
	configuration, err := config.Get()
	c, err := gorm.Open("postgres", configuration.GetString("database.postgres"))
	if err != nil {
		logger.Fatalf("couldn't connect to database: %v\n", err)
	}
	logger.Println("database connected")
	c.LogMode(true)
	c.SetLogger(logger)
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
	log.Printf("Database Initialization : '%s'", sql)
	c.Exec(sql)
	sql = "USE lpdp;"
	log.Printf("Database Initialization : '%s'", sql)
	c.Exec(sql)
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
		&models.ExitReason{},
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
	var admin models.User
	admin.Email = "admin@lpdp.org"
	admin.UserName = "admin"
	admin.Password = "password"
	admin.IsEmailVerified = true
	c.FirstOrCreate(&admin, models.User{
		UserName: "admin",
	})
	return
}

//Close ...
func (i Instance) Close() error {
	return i.c.Close()
}

// //GetQuery ...
// func GetQuery(params graphql.ResolveParams, items []interface{}) (interface{}, error) {
// 	var count int
// 	page := params.Args["page"].(int)
// 	logger.Printf("page requested : %d", page)
// 	itemsPerPage := params.Args["itemsPerPage"].(int)
// 	logger.Printf("number of items per page requested : %d", itemsPerPage)
// 	Inst.c.Model(items.(type)).Count(&count)
// 	pages := math.RoundToEven(float64(count) / float64(itemsPerPage))
// 	logger.Printf("computed total pages : %d", int(pages))
// 	offset := (page - 1) * itemsPerPage
// 	Inst.c.Offset(offset).Limit(itemsPerPage).Find(&items)
// 	return items, nil
// }
