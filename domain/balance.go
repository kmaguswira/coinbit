package domain

import "time"

type Balance struct {
	WalletID       string
	DepositAmount  []DepositAmount
	AboveThreshold bool
}

type DepositAmount struct {
	Value     float64
	Timestamp time.Time
}

func (t *Balance) GetTotal() float64 {
	total := float64(0)
	for _, amount := range t.DepositAmount {
		total += amount.Value
	}

	return total
}

func (t *Balance) IsAboveThreshold() {
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
