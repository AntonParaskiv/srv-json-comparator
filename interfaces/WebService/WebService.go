package WebService

import (
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonCompareInteractor/JsonCompareInteractorInterface"
)

type WebService struct {
	jsonCompareInteractor JsonCompareInteractorInterface.Interactor
}

func New() (ws *WebService) {
	ws = new(WebService)
	return
}

func (ws *WebService) SetJsonCompareInteractor(jsonCompareInteractor JsonCompareInteractorInterface.Interactor) *WebService {
	ws.jsonCompareInteractor = jsonCompareInteractor
	return ws
}
