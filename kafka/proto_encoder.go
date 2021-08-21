package kafka

import (
	"google.golang.org/protobuf/proto"
)

// ProtoEncoder encodes and decodes protobuf messages.
type ProtoEncoder struct {
	data []byte
	err  error
}

func NewProtoEncoder(message proto.Message) *ProtoEncoder {
	data, err := proto.Marshal(message)
	return &ProtoEncoder{
		data: data,
		err:  err,
	}
}

func (e *ProtoEncoder) Encode() ([]byte, error) {
	return e.data, e.err
}

func (*ProtoEncoder) Decode(data []byte, message proto.Message) error {
	return proto.Unmarshal(data, message)
}

func (e *ProtoEncoder) Length() int {
	return len(e.data)
}
