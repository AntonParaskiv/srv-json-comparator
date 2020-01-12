package JsonMarshaller

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type Marshaller struct {
}

func New() (m *Marshaller) {
	m = new(Marshaller)
	return
}

func (m *Marshaller) MarshalWriter(writer io.Writer, data interface{}) (err error) {
	err = json.NewEncoder(writer).Encode(data)
	if err != nil {
		err = errors.Errorf("marshal failed: %s", err.Error())
		return
	}
	return
}

func (m *Marshaller) UnMarshalReader(reader io.ReadCloser, data interface{}) (err error) {
	err = json.NewDecoder(reader).Decode(data)
	if err != nil {
		err = errors.Errorf("unMarshal failed: %s", err.Error())
		return
	}
	return
}

func (m *Marshaller) MarshalBytes(data interface{}) (marshalledData []byte, err error) {
	marshalledData, err = json.Marshal(data)
	if err != nil {
		err = errors.Errorf("marshal failed: %s", err.Error())
		return
	}
	return
}

func (m *Marshaller) UnMarshalBytes(marshalledData []byte, data interface{}) (err error) {
	err = json.Unmarshal(marshalledData, data)
	if err != nil {
		err = errors.Errorf("unMarshal failed: %s", err.Error())
		return
	}
	return
}
