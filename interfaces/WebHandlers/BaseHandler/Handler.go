package BaseHandler

import (
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebHandlers/HandlerInterface"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonCompareInteractor/JsonCompareInteractorInterface"
)

type WebHandler struct {
	jsonCompareInteractor JsonCompareInteractorInterface.Interactor
}

func New() (h *WebHandler) {
	h = new(WebHandler)
	return
}

func (h *WebHandler) SetJsonCompareInteractor(jsonCompareInteractor JsonCompareInteractorInterface.Interactor) *WebHandler {
	h.jsonCompareInteractor = jsonCompareInteractor
	return h
}

func (h *WebHandler) NewRequest() (request HandlerInterface.Request) {
	request = NewRequest()
	return
}

func (h *WebHandler) Handle(requestInterface HandlerInterface.Request) (response HandlerInterface.Response) {
	request := requestInterface.(*Request)

	result, err := h.jsonCompareInteractor.IsEqual(request.MethodParams.First, request.MethodParams.Second)
	if err != nil {
		// TODO: handle error
		return
	}

	response = NewResponse().SetIsEqual(result)
	return
}
