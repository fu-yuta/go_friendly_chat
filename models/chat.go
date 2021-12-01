package models

import (
	"errors"
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

func AddChat(c Chat) (*Chat, error) {
	if !db.NewRecord(c) {
		log.Println("NewRecord error")
		return nil, errors.New("NewRecord error")
	}
	err := db.Create(&c).Error
	if err != nil {
		log.Println("Create error")
		return nil, err
	}

	return &c, nil
}

func GetChat(id string) (*Chat, error) {
	var chat Chat

	db.First(&chat, id)
	if chat.Id == 0 {
		log.Println("not found chat error")
		return nil, errors.New("not found chat error")
	}

	return &chat, nil
}

func GetAllChats() []Chat {
	var allChats []Chat
	db.Find(&allChats)
	return allChats
}

func UpdateChat(id string, updateChat requests.UpdateChat) (*Chat, error) {
	var chat Chat
	db.First(&chat, id)
	if chat.Id == 0 {
		log.Println("not found chat error")
		return nil, errors.New("not found chat error")
	}

	chat.Message = updateChat.Message

	err := db.Save(&chat).Error
	if err != nil {
		log.Println("Update error")
		return nil, err
	}
	return &chat, nil
}

func DeleteChat(id string) (*Chat, error) {
	var chat Chat
	db.First(&chat, id)
	if chat.Id == 0 {
		log.Println("not found chat error")
		return nil, errors.New("not found chat error")
	}

	err := db.Delete(&chat).Error
	if err != nil {
		log.Println("Delete error")
		return nil, err
	}

	return &chat, nil
}
