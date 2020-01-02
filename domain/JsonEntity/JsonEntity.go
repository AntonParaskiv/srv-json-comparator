package JsonEntity

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonArray"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonObject"
)

type JsonEntity interface{}

// TODO: make test
func New() (je *JsonEntity) {
	je = new(JsonEntity)
	return
}

// TODO: make test
func NewFromInterface(entityInterface interface{}) (je JsonEntity) {
	je = new(JsonEntity)
	switch entityInterface.(type) {

	case []interface{}:
		array := JsonArray.JsonArray(entityInterface.([]interface{}))
		je = &array

	case map[string]interface{}:
		object := JsonObject.JsonObject(entityInterface.(map[string]interface{}))
		je = &object

	default:
		je = entityInterface
	}

	return
}
