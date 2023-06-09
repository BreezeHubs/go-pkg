package typexpkg

import "github.com/gogo/protobuf/proto"

func Proto(buf []byte, pb proto.Message) error {
	return proto.Unmarshal(buf, pb)
}
