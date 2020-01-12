package HandlerMock

type Response struct {
	MyResponseVar interface{} `json:"myResponseVar"`
}

func NewResponse() (r *Response) {
	r = new(Response)
	return
}

func (r *Response) SetMyResponseVar(myResponseVar interface{}) *Response {
	r.MyResponseVar = myResponseVar
	return r
}
