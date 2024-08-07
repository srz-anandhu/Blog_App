package api

import (
	"fmt"
	"net/http"
)

func StartServer(port string, h http.Handler) {
	if port[0] != ':' {
		port = ":" + port
		//fmt.Println("port is ", port)
	}
	fmt.Printf("server started on port = %s", port)
	http.ListenAndServe(port, h)
}
