package app

import (
	"blog/pkg/api"
)

func Start() {
	r := apiRouter()
	//fmt.Println("Server started at port: 8080....")
	api.StartServer(":8080", r)
}
