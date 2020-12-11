package proto

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var protoMarshalOption = protojson.MarshalOptions{UseProtoNames: true}

func MarshalJSON(m proto.Message) ([]byte, error) {
	return protoMarshalOption.Marshal(m)
}
