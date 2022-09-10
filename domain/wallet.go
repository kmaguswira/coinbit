package domain

import (
	"encoding/json"
	"time"

	messaging "github.com/lovoo/goka/examples/3-messaging"
)

type Wallet struct {
	WalletID       string
	DepositAmount  []DepositAmount
	AboveThreshold bool
}

type DepositAmount struct {
	Value     float64
	Timestamp time.Time
}

func (t *Wallet) GetTotal() float64 {
	total := float64(0)
	for _, amount := range t.DepositAmount {
		total += amount.Value
	}

	return total
}

func (t *Wallet) IsAboveThreshold() {
	if len(t.DepositAmount) > 0 && !t.AboveThreshold {
		total := float64(0)
		lastDeposit := t.DepositAmount[len(t.DepositAmount)-1]
		startTime := lastDeposit.Timestamp.Add(-2 * time.Minute)

		for _, DepositAmount := range t.DepositAmount {
			if DepositAmount.Timestamp.After(startTime) {
				total += DepositAmount.Value
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
	return json.Marshal(value)
}

func (t *Wallet) Decode(data []byte) (interface{}, error) {
	var m []messaging.Message
	err := json.Unmarshal(data, &m)
	return m, err
}

func (t *DepositAmount) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (t *DepositAmount) Decode(data []byte) (interface{}, error) {
	var m []messaging.Message
	err := json.Unmarshal(data, &m)
	return m, err
}
