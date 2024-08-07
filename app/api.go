package app

import (
	"blog/pkg/api"
)

func Start() {
	r := apiRouter()
	api.StartServer(":8080", r)
}
