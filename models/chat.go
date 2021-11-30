package models

import (
	"fmt"
	"go_friendly_chat/controllers/requests"
	"log"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	db = setupDB()
}

func setupDB() *gorm.DB {
	DBDriverName := beego.AppConfig.String("db_driver_name")
	DBName := beego.AppConfig.String("db_name")
	DBUserName := beego.AppConfig.String("db_user_name")
	DBUserPassword := beego.AppConfig.String("db_user_password")
	DBHost := beego.AppConfig.String("db_host")

	connectionInfo := fmt.Sprintf("%s:%s@%s/%s", DBUserName, DBUserPassword, DBHost, DBName)
	db, err := gorm.Open(DBDriverName, connectionInfo)

	if err != nil {
		log.Println(err.Error())
	}

	return db

}

type Chat struct {
	Id       uint   `gorm:"primary_key"`
	UserName string `gorm:"size:255"`
	Message  string `gorm:"size:255"`
}

func AddChat(c Chat) error {
	return nil
}

func GetChat(id string) (*Chat, error) {
	var chat Chat

	return &chat, nil
}

func GetAllChats() []Chat {
	var allChats []Chat
	return allChats
}

func UpdateChat(id string, updateChat requests.UpdateChat) error {
	return nil
}

func DeleteChat(id string) error {
	return nil
}
