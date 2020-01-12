package BaseHandler

type Response struct {
	IsEqual bool `json:"isEqual"`
}

func NewResponse() (r *Response) {
	r = new(Response)
	return
}

func (r *Response) SetIsEqual(isEqual bool) *Response {
	r.IsEqual = isEqual
	return r
}
