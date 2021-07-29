package presenter

import (
	"micheltank/cryptocurrency-data-service/internal/domain"
)

type BlockResponse struct {
	// The acronym of the network
	NetworkCode       string                `json:"networkCode"`
	// The height of the block in the blockchain, or its number
	BlockNumber       int                   `json:"blockNumber"`
	// The time at which this block was mined by the miner
	DateTime          string                `json:"dateTime"`
	// The block hash of the previous block in the blockchain
	PreviousBlockhash string                `json:"previousBlockhash"`
	// The block hash of the next block in the blockchain. NextBlockhash=null if this is the last block in the blockchain
	NextBlockhash     string                `json:"nextBlockhash"`
	// The size of the block in bytes
	Size              int                   `json:"size"`
	// The array of ids of all transactions in this block, starting with the newly generated coins (only the first 10)
	Transactions      []TransactionResponse `json:"transactions"`
}

func NewBlockResponse(block domain.Block) BlockResponse {
	return BlockResponse{
		NetworkCode:       string(block.GetNetworkCode()),
		BlockNumber:       block.GetBlockNumber(),
		DateTime:          block.GetDateTime().Format("02-01-2006 15:04"),
		PreviousBlockhash: block.GetPreviousBlockhash(),
		NextBlockhash:     block.GetNextBlockhash(),
		Size:              block.GetSize(),
		Transactions:      NewTransactionsResponse(block.GetTransactions()),
	}
}
