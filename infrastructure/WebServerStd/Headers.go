package WebServerStd

import "github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/WebServerInterface"

type Headers map[string][]string

func NewHeaders() *Headers {
	h := make(Headers)
	return &h
}

func (h *Headers) Add(key, value string) WebServerInterface.Headers {
	(*h)[key] = append((*h)[key], value)
	return h
}

func (h *Headers) Set(key, value string) WebServerInterface.Headers {
	(*h)[key] = []string{value}
	return h
}

func (h *Headers) Get(key string) (value string) {
	valueList, ok := (*h)[key]
	if !ok {
		return
	}

	if len(valueList) == 0 {
		return
	}

	value = valueList[0]
	return
}

func (h *Headers) Del(key string) WebServerInterface.Headers {
	delete(*h, key)
	return h
}
