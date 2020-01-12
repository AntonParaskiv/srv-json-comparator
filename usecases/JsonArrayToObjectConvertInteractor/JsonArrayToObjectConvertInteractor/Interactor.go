package JsonArrayToObjectConvertInteractor

import "github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/HashRepositoryInterface"

type Interactor struct {
	hashRepository HashRepositoryInterface.Repository
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}

func (i *Interactor) SetHashRepository(hashRepository HashRepositoryInterface.Repository) *Interactor {
	i.hashRepository = hashRepository
	return i
}
