package main

import "github.com/redshiftkeying/slowflow/router"

func main(){
	r := router.Router
	// Listen and Server in 0.0.0.0:8000 default
	router.SetRouter()

	r.Run(":8000")
}