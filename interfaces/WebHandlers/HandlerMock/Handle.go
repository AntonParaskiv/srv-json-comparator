package HandlerMock

import "github.com/AntonParaskiv/srv-json-comparator/interfaces/WebHandlers/HandlerInterface"

func (h *Handler) Handle(requestInterface HandlerInterface.Request) (response HandlerInterface.Response) {
	request := requestInterface.(*Request)
	h.SetRequest(request)
	response = h.GetResponse()
	return
}
