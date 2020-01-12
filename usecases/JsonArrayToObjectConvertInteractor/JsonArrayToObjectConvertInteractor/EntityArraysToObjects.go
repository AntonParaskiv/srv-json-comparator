package JsonArrayToObjectConvertInteractor

import (
	"fmt"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonArray"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonObject"
	"github.com/pkg/errors"
)

// TODO: cover tests
func (i *Interactor) EntityArraysToObjects(jsonEntityIn JsonEntity.JsonEntity) (jsonEntityOut JsonEntity.JsonEntity, err error) {
	switch jsonEntityIn.(type) {

	case *JsonArray.JsonArray:
		jsonArrayIn := jsonEntityIn.(*JsonArray.JsonArray)
		jsonEntityOut, err = i.convertArrayToObject(jsonArrayIn)
		if err != nil {
			err = errors.Errorf("convert array to object failed: %s", err.Error())
			return
		}

	case *JsonObject.JsonObject:
		jsonObjectIn := jsonEntityIn.(*JsonObject.JsonObject)
		jsonEntityOut, err = i.convertObjectToObject(jsonObjectIn)
		if err != nil {
			err = errors.Errorf("convert object to object failed: %s", err.Error())
			return
		}

	default:
		jsonEntityOut = jsonEntityIn
	}

	return
}

// TODO: cover tests
func (i *Interactor) convertArrayToObject(arrayIn *JsonArray.JsonArray) (objectOut *JsonObject.JsonObject, err error) {
	objectOut = JsonObject.New()

	for _, value := range *arrayIn {
		value = JsonEntity.NewFromInterface(value)
		value, err = i.EntityArraysToObjects(value)
		if err != nil {
			return
		}

		hash, err := i.hashRepository.GetHash(value)
		if err != nil {
			err = errors.Errorf("get hash of value %v failed", value)
			return nil, err
		}

		key := createFieldKeyForHash(objectOut, hash)
		objectOut.Add(key, value)
	}

	return
}

// TODO: cover tests
func (i *Interactor) convertObjectToObject(objectIn *JsonObject.JsonObject) (objectOut *JsonObject.JsonObject, err error) {
	objectOut = objectIn

	for key, value := range *objectIn {
		value = JsonEntity.NewFromInterface(value)
		value, err = i.EntityArraysToObjects(value)
		if err != nil {
			return
		}

		(*objectOut)[key] = value
	}

	return
}

// TODO: move function to object method
func createFieldKeyForHash(object *JsonObject.JsonObject, hash string) (key string) {
	for i := 1; true; i++ {
		key = fmt.Sprintf("%s_%d", hash, i)
		if object.IsKeyExist(key) {
			continue
		}
		break
	}
	return
}
