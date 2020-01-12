package WebServerStd

import "io"

type ResponseWriter struct {
	writer io.Writer
}

func NewResponseWriter() (r *ResponseWriter) {
	r = new(ResponseWriter)
	return
}

func (r *ResponseWriter) GetWriter() (writer io.Writer) {
	writer = r.writer
	return
}

func (r *ResponseWriter) SetWriter(writer io.Writer) *ResponseWriter {
	r.writer = writer
	return r
}
