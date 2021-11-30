package routers

import (
	"go_friendly_chat/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/chat",
			beego.NSInclude(
				&controllers.ChatController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
