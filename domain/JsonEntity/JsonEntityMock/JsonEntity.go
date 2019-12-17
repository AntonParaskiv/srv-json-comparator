package JsonEntityMock

type JsonEntity struct {
	data interface{}
}

func New() (je *JsonEntity) {
	je = new(JsonEntity)
	return
}
