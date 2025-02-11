package http201

import (
	"net/http"
)

func ServeLocation(responseWriter http.ResponseWriter, url string) error {
	if nil == responseWriter {
		return ErrNilHTTPResponseWriter
	}

	var header http.Header = responseWriter.Header()
	if nil == header {
		return errNilHTTPResponseWriterHeader
	}
	header.Add("Location", url)

	responseWriter.WriteHeader(StatusCode)
	return nil
}
