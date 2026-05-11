package http201

import (
	"io"
	"net/http"
)

func ServeString(responseWriter http.ResponseWriter, value string) error {
	if nil == responseWriter {
		return ErrNilHTTPResponseWriter
	}

	responseWriter.WriteHeader(StatusCode)

	_, err := io.WriteString(responseWriter, DefaultStatusText)
	if nil != err {
		return err
	}

	return nil
}
