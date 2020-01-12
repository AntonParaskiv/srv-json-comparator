package HandlerMock

type Request struct {
	MyRequestVar interface{} `json:"myRequestVar"`
}

func NewRequest() (r *Request) {
	r = new(Request)
	return
}

func (r *Request) SetMyRequestVar(myRequestVar interface{}) *Request {
	r.MyRequestVar = myRequestVar
	return r
}
