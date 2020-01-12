package WebService

import (
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebHandlers/HandlerInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/MarshallerInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/WebServerInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/WerServiceInterface"
)

type WebService struct {
	webServer  WebServerInterface.WebServer
	marshaller MarshallerInterface.Marshaller
}

func New() (ws *WebService) {
	ws = new(WebService)
	return
}

// TODO: cover test
func (ws *WebService) SetWebServer(webServer WebServerInterface.WebServer) WerServiceInterface.WebService {
	ws.webServer = webServer
	return ws
}

// TODO: cover test
func (ws *WebService) SetMarshaller(marshaller MarshallerInterface.Marshaller) WerServiceInterface.WebService {
	ws.marshaller = marshaller
	return ws
}

// TODO: cover test
func (ws *WebService) HandleFunc(path string, handler HandlerInterface.Handler) WerServiceInterface.WebService {
	serviceHandler := ws.createWebServiceHandler(handler)
	ws.webServer.HandleFunc(path, serviceHandler)
	return ws
}

// TODO: cover test
func (ws *WebService) createWebServiceHandler(handler HandlerInterface.Handler) (serviceHandler func(responseWriter WebServerInterface.ResponseWriter, request WebServerInterface.Request)) {
	serviceHandler = func(responseWriter WebServerInterface.ResponseWriter, request WebServerInterface.Request) {
		if request.Headers().Get("Content-Type") != "application/json" {
			// TODO: write error
			return
		}

		requestHandler := handler.NewRequest()

		err := ws.marshaller.UnMarshalReader(request.Body(), requestHandler)
		if err != nil {
			//http.Error(responseWriter.Writer(), err.Error(), http.StatusBadRequest)
			return
		}

		response := handler.Handle(requestHandler)
		err = ws.marshaller.MarshalWriter(responseWriter.Writer(), response)
		if err != nil {
			//http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}

		responseWriter.Headers().Add("Content-Type", "application/json")
	}
	return
}
