package WerServiceInterface

import (
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebHandlers/HandlerInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/MarshallerInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/WebServerInterface"
)

type WebService interface {
	SetWebServer(webServer WebServerInterface.WebServer) WebService
	SetMarshaller(marshaller MarshallerInterface.Marshaller) WebService
	HandleFunc(path string, handler HandlerInterface.Handler) WebService
}
