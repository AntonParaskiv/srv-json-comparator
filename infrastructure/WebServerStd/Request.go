package WebServerStd

import (
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/WebServerInterface"
	"io"
)

type Request struct {
	body    io.ReadCloser
	headers WebServerInterface.Headers
}

func NewRequest() (r *Request) {
	r = new(Request)
	return
}

func (r *Request) SetBody(body io.ReadCloser) *Request {
	r.body = body
	return r
}

func (r *Request) Body() (body io.ReadCloser) {
	body = r.body
	return
}

func (r *Request) SetHeaders(headers WebServerInterface.Headers) WebServerInterface.Request {
	r.headers = headers
	return r
}

func (r *Request) Headers() (headers WebServerInterface.Headers) {
	headers = r.headers
	return
}
