package proto

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func NewTimestamp(t time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
