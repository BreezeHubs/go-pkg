package rpcpkg

import "encoding/json"

type RpcResponse struct {
	Data []byte
	Err  error
}

func RpcReply(data any, err error) *RpcResponse {
	bytes, _ := json.Marshal(data)

	return &RpcResponse{
		Data: bytes,
		Err:  err,
	}
}
