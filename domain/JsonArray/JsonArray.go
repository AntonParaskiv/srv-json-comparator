package JsonArray

type JsonArray []interface{}

func New() (ja *JsonArray) {
	ja = &JsonArray{}
	return
}

func (ja *JsonArray) Add(value interface{}) *JsonArray {
	*ja = append(*ja, value)
	return ja
}
