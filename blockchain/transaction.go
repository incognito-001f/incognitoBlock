package blockchain

import "time"

type Transaction struct {
	Sender     []byte
	Receiver   []byte
	Amount     uint64
	Fee        uint64
	Size       int
	Priority   string
	Signature  []byte
	Timestamp  int64
	Commitment []byte
	Blinding   []byte
	KeyImage   []byte
}

func NewTransaction(sender, receiver []byte, amount uint64, size int, priority string) *Transaction {
	tx := &Transaction{
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Size:      size,
		Priority:  priority,
		Timestamp: time.Now().Unix(),
	}

	tx.Fee = tx.CalculateFee()
	return tx
}
