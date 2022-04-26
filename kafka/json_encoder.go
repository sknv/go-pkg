package kafka

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// JSONEncoder encodes and decodes JSON messages.
type JSONEncoder struct {
	data []byte
	err  error
}

func NewJSONEncoder(message any) *JSONEncoder {
	data, err := json.Marshal(message)
	return &JSONEncoder{
		data: data,
		err:  err,
	}
}

func (e *JSONEncoder) Encode() ([]byte, error) {
	return e.data, e.err
}

func (*JSONEncoder) Decode(data []byte, message any) error {
	return json.Unmarshal(data, message)
}

func (e *JSONEncoder) Length() int {
	return len(e.data)
}
