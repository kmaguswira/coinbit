package domain

import (
	"encoding/json"

	pb "github.com/kmaguswira/coinbit/proto"
)

type AboveThreshold struct {
	WalletID       string
	AboveThreshold bool
}

func (t *AboveThreshold) Encode(value interface{}) ([]byte, error) {
	aboveThresholds := value.([]AboveThreshold)

	aboveThresholdsProto := []pb.AboveThreshold{}
	for _, wallet := range aboveThresholds {
		aboveThresholdsProto = append(aboveThresholdsProto, pb.AboveThreshold{
			WalletID:       wallet.WalletID,
			AboveThreshold: wallet.AboveThreshold,
		})
	}

	return json.Marshal(&aboveThresholds)
}

func (t *AboveThreshold) Decode(data []byte) (interface{}, error) {
	var aboveThresholdsProto []pb.AboveThreshold
	err := json.Unmarshal(data, &aboveThresholdsProto)

	aboveThresholds := []AboveThreshold{}
	for _, aboveThreshold := range aboveThresholdsProto {
		aboveThresholds = append(aboveThresholds, AboveThreshold{
			WalletID:       aboveThreshold.WalletID,
			AboveThreshold: aboveThreshold.AboveThreshold,
		})
	}

	return aboveThresholds, err
}
