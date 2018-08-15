package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/jinzhu/gorm"
)

//DatabaseInitialization ...
func DatabaseInitialization(conn *gorm.DB) {

	if (!conn.HasTable(&models.Achievement{})) {
		conn.CreateTable(&models.Achievement{})
	}

	if (!conn.HasTable(&models.Activity{})) {
		conn.CreateTable(&models.Activity{})
	}

	return
}
