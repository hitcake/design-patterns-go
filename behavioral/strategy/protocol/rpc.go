package protocol

import "fmt"

type Rpc struct {
	call     func([]byte) []byte
	protocol Protocol
}

func (r *Rpc) SetProtocol(p Protocol) {
	r.protocol = p
}
func (r *Rpc) Request(param string) (string, error) {
	b, err := r.protocol.Encode(param)
	if err != nil {
		return "", err
	}
	resp := r.call(b)
	result, err := r.protocol.Decode(resp)
	if err != nil {
		return "", err
	}
	switch t := result.(type) {
	case string:
		return t, nil
	}
	return "", fmt.Errorf("unknow result")
}
