package protocol

import (
	"fmt"
	"testing"
)

func TestRpc_Request(t *testing.T) {
	rpc := &Rpc{call: func(b []byte) []byte {
		return b
	}, protocol: &Thrift{}}
	res, err := rpc.Request("hello")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
	// 更改策略，使用protobuf传输
	rpc.SetProtocol(&Protobuf{})
	res, err = rpc.Request("hello")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
