package JsonCompareInteractor

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"github.com/pkg/errors"
	"reflect"
)

func (i *Interactor) IsEqual(first, second JsonEntity.JsonEntity) (isEqual bool, err error) {
	var firstConverted, secondConverted JsonEntity.JsonEntity

	firstConverted, err = i.arrayToObjectConverter.EntityArraysToObjects(first)
	if err != nil {
		err = errors.Errorf("convert first entity failed: %s", err.Error())
		return
	}

	secondConverted, err = i.arrayToObjectConverter.EntityArraysToObjects(second)
	if err != nil {
		err = errors.Errorf("convert second entity failed: %s", err.Error())
		return
	}

	isEqual = reflect.DeepEqual(firstConverted, secondConverted)

	return
}
