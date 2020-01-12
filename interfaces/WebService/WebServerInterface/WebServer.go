package WebServerInterface

import (
	"io"
)

type WebServer interface {
	HandleFunc(pattern string, handler func(ResponseWriter, Request)) WebServer
	SetListenAddr(listenAddr string) WebServer
	Start() (err error)
}

type Request interface {
	Body() (body io.ReadCloser)
	Headers() (headers Headers)
}

type ResponseWriter interface {
	Writer() (writer io.Writer)
	Headers() (headers Headers)
}

type Headers interface {
	Add(key, value string) Headers
	Set(key, value string) Headers
	Get(key string) (value string)
	Del(key string) Headers
}

// TODO: write mock
