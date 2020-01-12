package HashRepositoryInterface

import "github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"

type Repository interface {
	GetHash(entity JsonEntity.JsonEntity) (hash string, err error)
}
