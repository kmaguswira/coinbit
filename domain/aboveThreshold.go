package domain

import (
	pb "github.com/kmaguswira/coinbit/proto"
	"google.golang.org/protobuf/proto"
)

type AboveThreshold struct {
	WalletID       string
	AboveThreshold bool
}

func (t *AboveThreshold) Encode(value interface{}) ([]byte, error) {
	aboveThreshold := value.(AboveThreshold)
	aboveThresholdProto := pb.AboveThreshold{
		WalletID:       aboveThreshold.WalletID,
		AboveThreshold: aboveThreshold.AboveThreshold,
	}

	return proto.Marshal(&aboveThresholdProto)
}

func (t *AboveThreshold) Decode(data []byte) (interface{}, error) {
	var aboveThresholdProto pb.AboveThreshold
	err := proto.Unmarshal(data, &aboveThresholdProto)

	aboveThreshold := AboveThreshold{
		WalletID:       aboveThresholdProto.WalletID,
		AboveThreshold: aboveThresholdProto.AboveThreshold,
	}

	return aboveThreshold, err
}
