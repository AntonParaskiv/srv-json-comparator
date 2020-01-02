package JsonMarshaller

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
)

type JsonMarshaller interface {
	Marshal(jsonEntity JsonEntity.JsonEntity) (jsonMarshaled []byte, err error)
	Unmarshal(jsonMarshaled []byte, jsonEntity JsonEntity.JsonEntity) (err error)
}
