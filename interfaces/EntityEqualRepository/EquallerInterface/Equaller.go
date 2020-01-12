package EquallerInterface

type Equaller interface {
	IsEqual(first, second interface{}) (isEqual bool)
}
