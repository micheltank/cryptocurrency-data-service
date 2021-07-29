package so_chain

type TransactionResponse struct {
	Status string                  `json:"status"`
	Data   TransactionDataResponse `json:"data"`
}

type TransactionDataResponse struct {
	TxId      string `json:"txid,omitempty"`
	Time      int64  `json:"time,omitempty"`
	SentValue string `json:"sent_value,omitempty"`
	Fee       string `json:"fee,omitempty"`
}