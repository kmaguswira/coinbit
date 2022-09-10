package domain

import (
	"time"

	pb "github.com/kmaguswira/coinbit/proto"
	"google.golang.org/protobuf/proto"
)

type Wallet struct {
	WalletID      string
	DepositAmount []DepositAmount
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

func (t *Wallet) Deposit(amount DepositAmount) {
	t.DepositAmount = append(t.DepositAmount, amount)
}

func (t *Wallet) IsAboveThreshold() {
	// if len(t.DepositAmount) > 0 && !t.AboveThreshold {
	// 	total := float64(0)
	// 	lastDeposit := t.DepositAmount[len(t.DepositAmount)-1]
	// 	startTime := lastDeposit.Timestamp.Add(-2 * time.Minute)

	// 	for _, deposit := range t.DepositAmount {
	// 		if deposit.Timestamp.After(startTime) {
	// 			total += deposit.Amount
	// 		}
	// 	}

	// 	if total > 10000 {
	// 		t.AboveThreshold = true
	// 		return
	// 	}
	// }

	// t.AboveThreshold = false
}

func (t *Wallet) Encode(value interface{}) ([]byte, error) {
	wallet := value.(Wallet)

	depositAmountProto := []*pb.DepositAmount{}
	for _, deposit := range wallet.DepositAmount {
		depositAmountProto = append(depositAmountProto, &pb.DepositAmount{
			Amount:    deposit.Amount,
			Timestamp: deposit.Timestamp.UnixMilli(),
		})
	}

	walletProto := pb.Wallet{
		WalletID:      wallet.WalletID,
		DepositAmount: depositAmountProto,
	}

	return proto.Marshal(&walletProto)
}

func (t *Wallet) Decode(data []byte) (interface{}, error) {
	var walletProto pb.Wallet
	err := proto.Unmarshal(data, &walletProto)

	depositAmount := []DepositAmount{}
	for _, deposit := range walletProto.DepositAmount {
		depositAmount = append(depositAmount, DepositAmount{
			Amount:    deposit.Amount,
			Timestamp: time.Unix(deposit.Timestamp, int64(time.Second)),
		})
	}

	wallet := Wallet{
		WalletID:      walletProto.WalletID,
		DepositAmount: depositAmount,
	}

	return wallet, err
}

func (t *DepositAmount) Encode(value interface{}) ([]byte, error) {
	depositAmount := value.(DepositAmount)
	depositAmountProto := pb.DepositAmount{
		Amount:    depositAmount.Amount,
		Timestamp: depositAmount.Timestamp.Unix(),
	}
	return proto.Marshal(&depositAmountProto)
}

func (t *DepositAmount) Decode(data []byte) (interface{}, error) {
	var depositAmountProto pb.DepositAmount
	err := proto.Unmarshal(data, &depositAmountProto)

	depositAmount := DepositAmount{
		Amount:    depositAmountProto.Amount,
		Timestamp: time.Unix(depositAmountProto.Timestamp, int64(time.Second)),
	}

	return depositAmount, err
}
