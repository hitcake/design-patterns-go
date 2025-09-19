package protocol

type Protocol interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}

type Protobuf struct {
}

func (p *Protobuf) Encode(msg interface{}) ([]byte, error) {
	return []byte("protobuf"), nil
}
func (p *Protobuf) Decode(data []byte) (interface{}, error) {
	return "protobuf", nil
}

type Thrift struct{}

func (t *Thrift) Encode(msg interface{}) ([]byte, error) {
	return []byte("thrift"), nil
}
func (t *Thrift) Decode(data []byte) (interface{}, error) {
	return "thrift", nil
}
