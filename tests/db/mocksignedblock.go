package db

import (
	"crypto/sha256"
	"encoding/binary"

	"github.com/coschain/contentos-go/common"
)

type MockSignedBlock struct {
	Payload []byte
	Num     uint64
	Prev    common.BlockID
}

func (msb *MockSignedBlock) Marshall() ([]byte, error) {
	return msb.Payload, nil
}

func (msb *MockSignedBlock) Unmarshall(b []byte) error {
	msb.Payload = b
	return nil
}

func (msb *MockSignedBlock) Set(data string, num uint64, prev common.BlockID) {
	msb.Payload = []byte(data)
	msb.Num = num
	msb.Prev = prev
}

func (msb *MockSignedBlock) Data() string {
	return string(msb.Payload)
}

func (msb *MockSignedBlock) Validate() bool {
	return true
}

func (msb *MockSignedBlock) GetSignee() (interface{}, error) {
	return nil, nil
}

func (msb *MockSignedBlock) Timestamp() uint64 {
	return 0
}

func (msb *MockSignedBlock) Id() common.BlockID {
	h := sha256.Sum256(msb.Payload)
	binary.LittleEndian.PutUint64(h[:8], msb.Num)
	return common.BlockID{
		Data: h,
	}
}

func (msb *MockSignedBlock) Previous() common.BlockID {
	return msb.Prev
}
