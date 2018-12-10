package models
import (
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)
type Templet struct {
	Id   int    `orm:"pk" json:"id"`
	Name string `orm:"size(100)" json:"name"`
	Type string   `orm:"size(1)" json:"type"`
}
func init(){
    orm.RegisterModel(new(Templet))
}