package WebServerStd

import "io"

type Request struct {
	body io.ReadCloser
}

func NewRequest() (r *Request) {
	r = new(Request)
	return
}

func (r *Request) GetBody() (body io.ReadCloser) {
	body = r.body
	return
}

func (r *Request) SetBody(body io.ReadCloser) *Request {
	r.body = body
	return r
}
