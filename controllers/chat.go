package controllers

import (
	"encoding/json"
	"go_friendly_chat/controllers/requests"
	"go_friendly_chat/controllers/responses"
	"go_friendly_chat/models"
	"log"

	"github.com/astaxie/beego"
)

type ChatController struct {
	beego.Controller
}

// @Title CreateChatMessage
// @Description create message
// @Param	body		body 	requests.Chat	true		"body for user content"
// @Success 200 {object} responses.Chat
// @Failure 403 : requests.Chat is empty
// @Failure 500 internal server error
// @router / [post]
func (c *ChatController) Post() {
	var chat requests.Chat

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &chat)
	if err != nil {
		log.Println("Chat Post json.Unmarshal error")
		c.Ctx.Output.SetStatus(403)
		c.ServeJSON()
	}

	newChat, err := models.AddChat(models.Chat{
		Id:       0,
		UserName: chat.UserName,
		Message:  chat.Message,
	})
	if err != nil {
		log.Println("AddChat error")
		c.Ctx.Output.SetStatus(500)
		c.ServeJSON()
	}

	res := responses.Chat{
		Id:       int(newChat.Id),
		UserName: newChat.UserName,
		Message:  newChat.Message,
	}
	c.Data["json"] = res

	c.ServeJSON()
}

// @Title GetAll
// @Description get all Chats
// @Success 200 {object} responses.Chats
// @router / [get]
func (c *ChatController) GetAll() {
	chats := models.GetAllChats()
	var res responses.Chats
	for _, chat := range chats {
		res.Chats = append(res.Chats, responses.Chat{
			Id:       int(chat.Id),
			UserName: chat.UserName,
			Message:  chat.Message,
		})

	}
	c.Data["json"] = res
	c.ServeJSON()
}

// @Title Get
// @Description get chat by id
// @Param	id		path 	string	true
// @Success 200 {object} responses.Chat
// @Failure 403 :id is empty
// @Failure 404 :chat is not found
// @router /:id [get]
func (c *ChatController) Get() {
	id := c.GetString(":id")
	if id != "" {
		chat, err := models.GetChat(id)
		if err != nil {
			log.Println("chat is not found")
			c.Ctx.Output.SetStatus(404)
			c.ServeJSON()
		} else {
			res := responses.Chat{
				Id:       int(chat.Id),
				UserName: chat.UserName,
				Message:  chat.Message,
			}
			c.Data["json"] = res
		}
		c.ServeJSON()
	} else {
		log.Println("id is empty error")
		c.Ctx.Output.SetStatus(403)
		c.ServeJSON()
	}
}

// @Title Update
// @Description update the chat
// @Param	id		path 	string	true
// @Param	body		body 	requests.UpdateChat	true
// @Success 200 {object} responses.Chat
// @Failure 403 :id is not int
// @Failure 500 internal server error
// @router /:id [put]
func (c *ChatController) Put() {
	id := c.GetString(":id")
	if id != "" {
		var req requests.UpdateChat
		json.Unmarshal(c.Ctx.Input.RequestBody, &req)
		updateChat, err := models.UpdateChat(id, req)
		if err != nil {
			c.Ctx.Output.SetStatus(500)
			c.ServeJSON()
		}
		res := responses.Chat{
			Id:       int(updateChat.Id),
			UserName: updateChat.UserName,
			Message:  updateChat.Message,
		}
		c.Data["json"] = res
		c.ServeJSON()
	} else {
		log.Println("id is empty error")
		c.Ctx.Output.SetStatus(403)
		c.ServeJSON()
	}
}

// @Title Delete
// @Description delete the chat
// @Param	id		path 	string	true
// @Success 200 {object} responses.Chat
// @Failure 403 id is empty
// @Failure 500 internal server error
// @router /:id [delete]
func (c *ChatController) Delete() {
	id := c.GetString(":id")
	if id != "" {
		deleteChat, err := models.DeleteChat(id)
		if err != nil {
			log.Println("Delete error")
			c.Ctx.Output.SetStatus(500)
			c.ServeJSON()
		}
		res := responses.Chat{
			Id:       int(deleteChat.Id),
			UserName: deleteChat.UserName,
			Message:  deleteChat.Message,
		}
		c.Data["json"] = res
		c.ServeJSON()
	} else {
		log.Println("id is empty error")
		c.Ctx.Output.SetStatus(403)
		c.ServeJSON()
	}
}
