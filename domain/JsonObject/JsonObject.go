package JsonObject

type JsonObject map[string]interface{}

func New() (jo *JsonObject) {
	jo = &JsonObject{}
	return
}

func (jo *JsonObject) Add(key string, value interface{}) *JsonObject {
	(*jo)[key] = value
	return jo
}

func (jo *JsonObject) IsKeyExist(key string) (isExist bool) {
	_, isExist = (*jo)[key]
	return
}
