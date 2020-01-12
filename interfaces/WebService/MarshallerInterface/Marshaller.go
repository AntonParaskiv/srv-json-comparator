package MarshallerInterface

import "io"

type Marshaller interface {
	MarshalWriter(writer io.Writer, data interface{}) (err error)
	UnMarshalReader(reader io.ReadCloser, data interface{}) (err error)
}
