package JsonArrayMock

type JsonArray []interface{}

func New() (ja *JsonArray) {
	ja = &JsonArray{}
	return
}
