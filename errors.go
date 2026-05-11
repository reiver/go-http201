package http201

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrNilHTTPResponseWriter = erorr.Error("http201: nil http-response-writer")
)

const (
	errNilHTTPResponseWriterHeader = erorr.Error("http201: nil http-response-writer header")
)
