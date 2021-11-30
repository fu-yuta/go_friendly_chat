package models

import (
	"go_friendly_chat/controllers/requests"
)

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
