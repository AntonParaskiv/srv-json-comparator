package EquallerMock

type Equaller struct {
	first   interface{}
	second  interface{}
	isEqual bool
}

func New() (e *Equaller) {
	e = new(Equaller)
	return
}

func (e *Equaller) SetFirst(first interface{}) *Equaller {
	e.first = first
	return e
}

func (e *Equaller) GetFirst() (first interface{}) {
	first = e.first
	return
}

func (e *Equaller) SetSecond(second interface{}) *Equaller {
	e.second = second
	return e
}

func (e *Equaller) GetSecond() (second interface{}) {
	second = e.second
	return
}

func (e *Equaller) SetIsEqual(isEqual bool) *Equaller {
	e.isEqual = isEqual
	return e
}

func (e *Equaller) GetIsEqual() (isEqual bool) {
	isEqual = e.isEqual
	return
}

func (e *Equaller) IsEqual(first, second interface{}) (isEqual bool) {
	e.SetFirst(first)
	e.SetSecond(second)
	isEqual = e.GetIsEqual()
	return
}
