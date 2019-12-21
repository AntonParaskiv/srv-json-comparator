package JsonObjectMock

type JsonObject map[string]interface{}

func New() (jo *JsonObject) {
	jo = &JsonObject{}
	return
}
