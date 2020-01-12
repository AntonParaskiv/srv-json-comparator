package BaseHandler

type Request struct {
	Method       string `json:"method,omitempty"`
	MethodParams struct {
		First  interface{} `json:"first,omitempty"`
		Second interface{} `json:"second,omitempty"`
	} `json:"methodParams,omitempty"`
}

func NewRequest() (r *Request) {
	r = new(Request)
	return
}
