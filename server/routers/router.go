package routers

import (
	"server/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.MainController{})
	beego.Router("/templets", &controllers.TempletController{}, "Get:GetTemplets")
}
