package JsonEntityMock

type JsonEntity struct {
	data interface{}
}

func New() (je *JsonEntity) {
	je = new(JsonEntity)
	return
}

func (je *JsonEntity) SetData(data interface{}) *JsonEntity {
	je.data = data
	return je
}
