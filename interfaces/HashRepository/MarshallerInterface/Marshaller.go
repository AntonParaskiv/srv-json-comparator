package MarshallerInterface

type Marshaller interface {
	MarshalBytes(data interface{}) (marshalledData []byte, err error)
	UnMarshalBytes(marshalledData []byte, data interface{}) (err error)
}
