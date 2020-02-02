package MarshallerMock

import (
	"io"
)

type Marshaller struct {
	writer         io.Writer
	reader         io.ReadCloser
	data           interface{}
	marshalledData []byte

	simulateErrorStepMatch int
	simulateErrorFlag      bool
}

func New() (m *Marshaller) {
	m = new(Marshaller)
	return
}

func (m *Marshaller) SetWriter(writer io.Writer) *Marshaller {
	m.writer = writer
	return m
}

func (m *Marshaller) GetWriter() (writer io.Writer) {
	writer = m.writer
	return
}

func (m *Marshaller) SetReader(reader io.ReadCloser) *Marshaller {
	m.reader = reader
	return m
}

func (m *Marshaller) GetReader() (reader io.ReadCloser) {
	reader = m.reader
	return
}

func (m *Marshaller) SetData(data interface{}) *Marshaller {
	m.data = data
	return m
}

func (m *Marshaller) GetData() (data interface{}) {
	data = m.data
	return
}

func (m *Marshaller) SetMarshalledData(marshalledData []byte) *Marshaller {
	m.marshalledData = marshalledData
	return m
}

func (m *Marshaller) GetMarshalledData() (marshalledData []byte) {
	marshalledData = m.marshalledData
	return
}

func (m *Marshaller) MarshalWriter(writer io.Writer, data interface{}) (err error) {
	m.SetWriter(writer)
	m.SetData(data)

	if m.IsSetSimulateError() {
		err = m.Error()
	}

	return
}

func (m *Marshaller) UnMarshalReader(reader io.ReadCloser, data interface{}) (err error) {
	m.SetReader(reader)
	m.SetData(data)

	if m.IsSetSimulateError() {
		err = m.Error()
	}

	return
}

func (m *Marshaller) MarshalBytes(data interface{}) (marshalledData []byte, err error) {
	m.SetData(data)

	marshalledData = m.GetMarshalledData()
	if m.IsSetSimulateError() {
		err = m.Error()
	}
	return
}

func (m *Marshaller) UnMarshalBytes(marshalledData []byte, data interface{}) (err error) {
	m.SetMarshalledData(marshalledData)
	m.SetData(data)

	if m.IsSetSimulateError() {
		err = m.Error()
	}

	return
}
