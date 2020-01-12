package JsonCompareInteractor

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"github.com/pkg/errors"
)

// TODO: cover tests
func (i *Interactor) IsEqual(first, second JsonEntity.JsonEntity) (isEqual bool, err error) {
	var firstConverted, secondConverted JsonEntity.JsonEntity

	// TODO: move to method
	firstEntity := JsonEntity.NewFromInterface(first)
	firstConverted, err = i.arrayToObjectConverter.EntityArraysToObjects(firstEntity)
	if err != nil {
		err = errors.Errorf("convert first entity failed: %s", err.Error())
		return
	}

	// TODO: move to method
	secondEntity := JsonEntity.NewFromInterface(second)
	secondConverted, err = i.arrayToObjectConverter.EntityArraysToObjects(secondEntity)
	if err != nil {
		err = errors.Errorf("convert second entity failed: %s", err.Error())
		return
	}

	// TODO: move to method
	isEqual = i.entityEqualRepository.IsEqual(firstConverted, secondConverted)

	return
}
