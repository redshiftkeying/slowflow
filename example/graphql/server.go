package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	sw "github.com/redshiftkeying/slowflow-server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//use default configs
	sw.SetServerOptions("")
	// Instance
	s := sw.Init()
	// add Middleware
	s.AddServerMiddleware(filter)
	// start server
	s.StartServer()
}

func filter(c *gin.Context) {
	fmt.Printf("请求参数：%s \n", c.FullPath())
	c.Next()
}
