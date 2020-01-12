package HandlerMock

import "github.com/AntonParaskiv/srv-json-comparator/interfaces/WebHandlers/HandlerInterface"

type Handler struct {
	request  *Request
	response *Response
}

func (h *Handler) SetRequest(request *Request) *Handler {
	h.request = request
	return h
}

func (h *Handler) GetRequest() (request *Request) {
	request = h.request
	return
}

func (h *Handler) SetResponse(response *Response) *Handler {
	h.response = response
	return h
}

func (h *Handler) GetResponse() (response *Response) {
	response = h.response
	return
}

func (h *Handler) NewRequest() (request HandlerInterface.Request) {
	request = NewRequest()
	return
}
