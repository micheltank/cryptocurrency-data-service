package so_chain

type BlockResponse struct {
	Status string            `json:"status"`
	Data   BlockDataResponse `json:"data"`
}

type BlockDataResponse struct {
	Network           string   `json:"network,omitempty"`
	Blockhash         string   `json:"blockhash,omitempty"`
	BlockNo           int      `json:"block_no,omitempty"`
	MiningDifficulty  string   `json:"mining_difficulty,omitempty"`
	Time              int64    `json:"time"`
	Confirmations     int      `json:"confirmations,omitempty"`
	IsOrphan          bool     `json:"is_orphan,omitempty"`
	Txs               []string `json:"txs,omitempty"`
	Merkleroot        string   `json:"merkleroot,omitempty"`
	PreviousBlockhash string   `json:"previous_blockhash,omitempty"`
	NextBlockhash     string   `json:"next_blockhash,omitempty"`
	Size              int      `json:"size,omitempty"`
}