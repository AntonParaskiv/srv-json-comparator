package HandlerInterface

type Handler interface {
	Handle(request Request) (response Response)
	NewRequest() (request Request)
}
