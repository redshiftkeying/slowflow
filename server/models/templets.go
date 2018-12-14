package models
import (
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
    "fmt"
)

type Templet struct {
	Id   int    `orm:"pk" json:"id"`
	Name string `orm:"size(100)" json:"name"`
	Type string   `orm:"size(1)" json:"type"`
}

type Templets struct {
	Templets [] Templet `json: "iterms"`
}

func init(){
    orm.RegisterModel(new(Templet))
}

func GetTemplets() Templets {
	sql := "select * from templet"

	o := orm.NewOrm()
	o.Using("default")
	
	var maps []orm.Params
	if _, err := o.Raw(sql).Values(&maps); err != nil {
		log.Fatalf("query templet failed that error was: %s", err.Error())
	}

	result := Templets{}
	for _, templet := range maps {
        fmt.Println(templet)
		id, _ := strconv.Atoi(templet["id"].(string))
		data := Templet{
			Id:   id,
			Name: templet["name"].(string),
			// Type: type,
		}
		result.Templets = append(result.Templets, data)
	}
	return result
}