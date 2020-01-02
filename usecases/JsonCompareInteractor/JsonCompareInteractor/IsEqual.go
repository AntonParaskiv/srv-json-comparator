package JsonCompareInteractor

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"github.com/pkg/errors"
	"reflect"
)

// TODO: cover tests
func (i *Interactor) IsEqual(first, second JsonEntity.JsonEntity) (isEqual bool, err error) {
	var firstConverted, secondConverted JsonEntity.JsonEntity

	// TODO: move to method
	firstConverted, err = i.arrayToObjectConverter.EntityArraysToObjects(first)
	if err != nil {
		err = errors.Errorf("convert first entity failed: %s", err.Error())
		return
	}

	// TODO: move to method
	secondConverted, err = i.arrayToObjectConverter.EntityArraysToObjects(second)
	if err != nil {
		err = errors.Errorf("convert second entity failed: %s", err.Error())
		return
	}

	// TODO: move to method
	isEqual = reflect.DeepEqual(firstConverted, secondConverted)

	return
}
