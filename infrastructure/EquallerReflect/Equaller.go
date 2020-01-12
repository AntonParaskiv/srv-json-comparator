package EquallerReflect

import "reflect"

type Equaller struct {
}

func New() (e *Equaller) {
	e = new(Equaller)
	return
}

func (e *Equaller) IsEqual(first, second interface{}) (isEqual bool) {
	isEqual = reflect.DeepEqual(first, second)
	return
}
