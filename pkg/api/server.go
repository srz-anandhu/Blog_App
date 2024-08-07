package api

import (
	"net/http"
)

func StartServer(port string, h http.Handler) {

	http.ListenAndServe(port, h)
}
