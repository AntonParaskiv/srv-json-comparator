package WebServerStd

import (
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/WebServerInterface"
	"io"
)

type ResponseWriter struct {
	writer  io.Writer
	headers WebServerInterface.Headers
}

func NewResponseWriter() (r *ResponseWriter) {
	r = new(ResponseWriter)
	return
}

func (r *ResponseWriter) SetWriter(writer io.Writer) *ResponseWriter {
	r.writer = writer
	return r
}

func (r *ResponseWriter) Writer() (writer io.Writer) {
	writer = r.writer
	return
}

func (r *ResponseWriter) SetHeaders(headers WebServerInterface.Headers) WebServerInterface.ResponseWriter {
	r.headers = headers
	return r
}

func (r *ResponseWriter) Headers() (headers WebServerInterface.Headers) {
	headers = r.headers
	return
}
