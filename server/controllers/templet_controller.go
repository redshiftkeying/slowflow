package controllers

import (
	// "encoding/json"
	"server/models"
	"github.com/astaxie/beego"
)

type H map[string]interface{}

type TempletController struct {
	beego.Controller
}

func (t *TempletController) GetTemplets() {
	datas := models.GetTemplets()
	t.Data["json"] = datas
	t.ServeJSON()
}
