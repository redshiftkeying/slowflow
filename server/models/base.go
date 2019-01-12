package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
  "fmt"
)

var engine *xorm.Engine

func init() {
  var err error
  engine, err = xorm.NewEngine("mysql", "root:8*PKKNP~y`Ft&TZIiTL-@/slowflows?charset=utf8")
  if err != nil {
    fmt.Println(err)
  }
}
