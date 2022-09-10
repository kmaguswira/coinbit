package domain

import (
	"encoding/json"
	"time"

	pb "github.com/kmaguswira/coinbit/proto"
	"google.golang.org/protobuf/proto"
)

type Wallet struct {
	WalletID       string
	DepositAmount  []DepositAmount
	AboveThreshold bool
}

type DepositAmount struct {
	Amount    float64
	Timestamp time.Time
}

func (t *Wallet) GetTotal() float64 {
	total := float64(0)
	for _, deposit := range t.DepositAmount {
		total += deposit.Amount
	}

	return total
}

func (t *Wallet) IsAboveThreshold() {
	if len(t.DepositAmount) > 0 && !t.AboveThreshold {
		total := float64(0)
		lastDeposit := t.DepositAmount[len(t.DepositAmount)-1]
		startTime := lastDeposit.Timestamp.Add(-2 * time.Minute)

		for _, deposit := range t.DepositAmount {
			if deposit.Timestamp.After(startTime) {
				total += deposit.Amount
			}
		}

		if total > 10000 {
			t.AboveThreshold = true
			return
		}
	}

	t.AboveThreshold = false
}

func (t *Wallet) Encode(value interface{}) ([]byte, error) {
	wallets := value.([]pb.Wallet)
	return json.Marshal(&wallets)
}

func (t *Wallet) Decode(data []byte) (interface{}, error) {
	var event []pb.Wallet
	err := json.Unmarshal(data, &event)
	return event, err
}

func (t *DepositAmount) Encode(value interface{}) ([]byte, error) {
	depositAmount := value.(pb.DepositAmount)
	return proto.Marshal(&depositAmount)
}

func (t *DepositAmount) Decode(data []byte) (interface{}, error) {
	event := &pb.DepositAmount{}
	err := proto.Unmarshal(data, event)
	return event, err
}
