package JsonCompareInteractor

import (
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/JsonArrayToObjectConvertInteractorInterface"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonCompareInteractor/EntityEqualRepositoryInterface"
)

type Interactor struct {
	arrayToObjectConverter JsonArrayToObjectConvertInteractorInterface.Interactor
	entityEqualRepository  EntityEqualRepositoryInterface.Repository
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}

func (i *Interactor) SetArrayToObjectConverter(arrayToObjectConverter JsonArrayToObjectConvertInteractorInterface.Interactor) *Interactor {
	i.arrayToObjectConverter = arrayToObjectConverter
	return i
}

func (i *Interactor) SetEntityEqualRepository(entityEqualRepository EntityEqualRepositoryInterface.Repository) *Interactor {
	i.entityEqualRepository = entityEqualRepository
	return i
}
