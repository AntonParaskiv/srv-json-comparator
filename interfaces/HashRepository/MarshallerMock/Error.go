package MarshallerMock

import (
	"fmt"
)

func (m *Marshaller) SimulateError(stepMatch int) *Marshaller {
	m.simulateErrorStepMatch = stepMatch
	m.simulateErrorFlag = true
	return m
}

func (m *Marshaller) IsSetSimulateError() (isSetSimulateError bool) {
	if !m.simulateErrorFlag {
		return
	}

	if m.simulateErrorStepMatch > 0 {
		m.simulateErrorStepMatch--
		return
	}

	isSetSimulateError = true
	return
}

func (m *Marshaller) Error() (err error) {
	m.simulateErrorFlag = false
	err = fmt.Errorf("simulated error")
	return
}
