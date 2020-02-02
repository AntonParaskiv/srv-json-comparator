package JsonCompareInteractorMock

import (
	"fmt"
)

func (i *Interactor) SimulateError(stepMatch int) *Interactor {
	i.simulateErrorStepMatch = stepMatch
	i.simulateErrorFlag = true
	return i
}

func (i *Interactor) IsSetSimulateError() (isSetSimulateError bool) {
	if !i.simulateErrorFlag {
		return
	}

	if i.simulateErrorStepMatch > 0 {
		i.simulateErrorStepMatch--
		return
	}

	isSetSimulateError = true
	return
}

func (i *Interactor) Error() (err error) {
	i.simulateErrorFlag = false
	err = fmt.Errorf("simulated error")
	return
}
