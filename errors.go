package http201

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrNilHTTPResponseWriter = erorr.Error("http201: nil http-response-writer")
)

const (
	errNilHTTPResponseWriterHeader = erorr.Error("http201: nil http-response-writer header")
)
