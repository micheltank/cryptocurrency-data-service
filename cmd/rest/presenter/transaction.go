package presenter

import (
	"micheltank/cryptocurrency-data-service/internal/domain"
)

type TransactionResponse struct {
	// The transaction id
	TransactionId string  `json:"transactionId"`
	// The time at which this transaction received by SoChain, or was mined by the miner
	DateTime      string  `json:"dateTime"`
	// The fee paid to the miner
	Fee           float64 `json:"fee"`
	// The total value of all coins sent in this transaction
	SentValue     float64 `json:"sentValue"`
}

func NewTransactionResponse(transaction domain.Transaction) TransactionResponse {
	return TransactionResponse{
		TransactionId: transaction.GetTransactionId(),
		DateTime:      transaction.GetDateTime().Format("02-01-2006 15:04"),
		Fee:           transaction.GetFee(),
		SentValue:     transaction.GetSentValue(),
	}
}

func NewTransactionsResponse(transactions domain.Transactions) (response []TransactionResponse) {
	for _, transaction := range transactions {
		response = append(response, NewTransactionResponse(transaction))
	}
	return response
}
