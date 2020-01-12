package WebServerStd

import (
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/WebServerInterface"
	"net/http"
)

// TODO: cover all test

type WebServer struct {
	muxServer  *http.ServeMux
	listenAddr string
}

const DefaultAddr = ":8080"

func New() (ws *WebServer) {
	ws = new(WebServer)
	ws.SetListenAddr(DefaultAddr)
	ws.SetServer(newMuxServer())
	return
}

func (ws *WebServer) SetListenAddr(listenAddr string) WebServerInterface.WebServer {
	ws.listenAddr = listenAddr
	return ws
}

func (ws *WebServer) GetListenAddr() (addr string) {
	return ws.listenAddr
}

func (ws *WebServer) SetServer(muxServer *http.ServeMux) *WebServer {
	ws.muxServer = muxServer
	return ws
}

func newMuxServer() (muxServer *http.ServeMux) {
	muxServer = http.NewServeMux()
	return
}

func (ws *WebServer) HandleFunc(pattern string, serviceHandler func(WebServerInterface.ResponseWriter, WebServerInterface.Request)) WebServerInterface.WebServer {
	serverHandler := createServerHandler(serviceHandler)
	ws.muxServer.HandleFunc(pattern, serverHandler)
	return ws
}

func createServerHandler(serviceHandler func(WebServerInterface.ResponseWriter, WebServerInterface.Request)) (serverHandler func(http.ResponseWriter, *http.Request)) {
	serverHandler = func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		responseWriter := NewResponseWriter().SetWriter(httpResponseWriter)
		request := NewRequest().SetBody(httpRequest.Body)
		serviceHandler(responseWriter, request)
	}
	return
}

func (ws *WebServer) Start() (err error) {
	err = http.ListenAndServe(ws.listenAddr, ws.muxServer)
	return
}
