package domain

import "time"

type Transaction struct {
	transactionId string
	dateTime      time.Time
	fee           float64
	sentValue     float64
}

type Transactions []Transaction

func NewTransaction(transactionId string,
	dateTime time.Time,
	fee float64,
	sentValue float64) Transaction {
	return Transaction{
		transactionId: transactionId,
		dateTime:      dateTime,
		fee:           fee,
		sentValue:     sentValue,
	}
}

func (t *Transaction) GetTransactionId() string {
	return t.transactionId
}

func (t *Transaction) GetDateTime() time.Time {
	return t.dateTime
}

func (t *Transaction) GetFee() float64 {
	return t.fee
}

func (t *Transaction) GetSentValue() float64 {
	return t.sentValue
}
