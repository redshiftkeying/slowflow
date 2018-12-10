package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "user-zKKrHv6txr:8*PKKNP~y`Ft&TZIiTL-@tcp(mysql.coding.io:3306)/db-JoeXsz4uC5?charset=utf8")
}