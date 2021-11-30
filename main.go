package main

import (
	"fmt"
	"go_friendly_chat/models"
	_ "go_friendly_chat/routers"
	"log"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

func main() {
	db := setupDB()
	db.AutoMigrate(&models.Chat{})

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func setupDB() *gorm.DB {
	dbDriverName := beego.AppConfig.String("db_driver_name")
	dbName := beego.AppConfig.String("db_name")
	dbUserName := beego.AppConfig.String("db_user_name")
	dbUserPassword := beego.AppConfig.String("db_user_password")
	dbHost := beego.AppConfig.String("db_host")

	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, dbUserName, dbUserPassword, dbHost, dbName)
	db, err := gorm.Open(dbDriverName, connect)

	if err != nil {
		log.Println(err.Error())
	}

	return db

}
