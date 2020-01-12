package EntityEqualRepositoryInterface

import "github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"

type Repository interface {
	IsEqual(first, second JsonEntity.JsonEntity) (isEqual bool)
}
