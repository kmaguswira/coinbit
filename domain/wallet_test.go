package domain

import (
	"fmt"
	"testing"
	"time"
)

func TestWalletDeposit(t *testing.T) {
	wallet := Wallet{
		WalletID: "1",
		DepositAmount: []DepositAmount{
			{
				Amount:    1000,
				Timestamp: time.Now(),
			},
		},
	}

	wallet.Deposit(DepositAmount{
		Amount:    500,
		Timestamp: time.Now(),
	})

	if len(wallet.DepositAmount) != 2 {
		t.Errorf("should be add deposit amount")
	}

	if wallet.DepositAmount[len(wallet.DepositAmount)-1].Amount != 500 {
		t.Errorf("last deposit should be 500")
	}
}

func TestWalletGetBalance(t *testing.T) {
	wallet := Wallet{
		WalletID: "1",
		DepositAmount: []DepositAmount{
			{
				Amount:    1000,
				Timestamp: time.Now(),
			},
			{
				Amount:    500,
				Timestamp: time.Now(),
			},
		},
	}

	totalBalance := wallet.GetTotal()

	if totalBalance != 1500 {
		t.Errorf("should be sum of the all deposit amounts")
	}
}

func TestWalletIsAboveThresholdLessThan10000(t *testing.T) {
	wallet := Wallet{
		WalletID: "1",
		DepositAmount: []DepositAmount{
			{
				Amount:    5000,
				Timestamp: time.Date(2022, 9, 10, 22, 0, 0, 0, time.UTC),
			},
			{
				Amount:    5000,
				Timestamp: time.Date(2022, 9, 10, 22, 1, 0, 0, time.UTC),
			},
		},
	}

	IsAboveThreshold := wallet.IsAboveThreshold()

	if IsAboveThreshold != false {
		t.Errorf("should be false cause total is less than 10000")
	}
}
func TestWalletIsAboveThresholdOK(t *testing.T) {
	wallet := Wallet{
		WalletID: "1",
		DepositAmount: []DepositAmount{
			{
				Amount:    6000,
				Timestamp: time.Date(2022, 9, 10, 22, 0, 0, 0, time.UTC),
			},
			{
				Amount:    6000,
				Timestamp: time.Date(2022, 9, 10, 22, 1, 0, 0, time.UTC),
			},
		},
	}

	IsAboveThreshold := wallet.IsAboveThreshold()

	if IsAboveThreshold != true {
		t.Errorf("should be true cause amount is greater than 10000 in span time 2 minutes")
	}
}

func TestWalletIsAboveThresholdDepositAfter2Minutes(t *testing.T) {
	wallet := Wallet{
		WalletID: "1",
		DepositAmount: []DepositAmount{
			{
				Amount:    6000,
				Timestamp: time.Date(2022, 9, 10, 22, 0, 0, 0, time.UTC),
			},
			{
				Amount:    6000,
				Timestamp: time.Date(2022, 9, 10, 22, 2, 0, 0, time.UTC),
			},
		},
	}

	IsAboveThreshold := wallet.IsAboveThreshold()

	fmt.Println(IsAboveThreshold)

	if IsAboveThreshold != false {
		t.Errorf("should be false cause not in span time 2 minutes")
	}
}

func TestWalletIsAboveThresholdEqual10000(t *testing.T) {
	wallet := Wallet{
		WalletID: "1",
		DepositAmount: []DepositAmount{
			{
				Amount:    5000,
				Timestamp: time.Date(2022, 9, 10, 22, 0, 0, 0, time.UTC),
			},
			{
				Amount:    5001,
				Timestamp: time.Date(2022, 9, 10, 22, 2, 0, 0, time.UTC),
			},
		},
	}

	IsAboveThreshold := wallet.IsAboveThreshold()

	if IsAboveThreshold != false {
		t.Errorf("should be false cause amount is equal 10000")
	}
}

func TestWalletIsAboveThreshold5Times2000In2MinutesAnd6000After5sec(t *testing.T) {
	wallet := Wallet{
		WalletID: "1",
		DepositAmount: []DepositAmount{
			{
				Amount:    2000,
				Timestamp: time.Date(2022, 9, 10, 22, 0, 0, 0, time.UTC),
			},
			{
				Amount:    2000,
				Timestamp: time.Date(2022, 9, 10, 22, 0, 24, 0, time.UTC),
			},
			{
				Amount:    2000,
				Timestamp: time.Date(2022, 9, 10, 22, 0, 48, 0, time.UTC),
			},
			{
				Amount:    2000,
				Timestamp: time.Date(2022, 9, 10, 22, 1, 12, 0, time.UTC),
			},
			{
				Amount:    2000,
				Timestamp: time.Date(2022, 9, 10, 22, 1, 36, 0, time.UTC),
			},
			{
				Amount:    6000,
				Timestamp: time.Date(2022, 9, 10, 22, 2, 5, 0, time.UTC),
			},
		},
	}

	IsAboveThreshold := wallet.IsAboveThreshold()

	if IsAboveThreshold != true {
		t.Errorf("should be true cause amount is greater than 10000 in time span 2 minutes")
	}
}
