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
	GetBody() (body io.ReadCloser)
}

type ResponseWriter interface {
	GetWriter() (writer io.Writer)
}

// TODO: write mock
