package routers

import (
	"server/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.MainController{})
	beego.Router("/temples", &controllers.TempletController{}, "Get:GetTemplets")
}
