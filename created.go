package http201

import (
	"net/http"
)

func Created(responseWriter http.ResponseWriter, request *http.Request) {
	Serve(responseWriter)
}
