package domain

import "time"

type Block struct {
	networkCode       NetworkCode
	blockNumber       int
	dateTime          time.Time
	previousBlockhash string
	nextBlockhash     string
	size              int
	transactions      []Transaction
}

func NewBlock(networkCode NetworkCode,
	blockNumber int,
	dateTime time.Time,
	previousBlockhash string,
	nextBlockhash string,
	size int,
	transactions []Transaction) Block {
	return Block{
		networkCode:       networkCode,
		blockNumber:       blockNumber,
		dateTime:          dateTime,
		previousBlockhash: previousBlockhash,
		nextBlockhash:     nextBlockhash,
		size:              size,
		transactions:      transactions,
	}
}

func (b *Block) GetNetworkCode() NetworkCode {
	return b.networkCode
}

func (b *Block) GetBlockNumber() int {
	return b.blockNumber
}

func (b *Block) GetDateTime() time.Time {
	return b.dateTime
}

func (b *Block) GetPreviousBlockhash() string {
	return b.previousBlockhash
}

func (b *Block) GetNextBlockhash() string {
	return b.nextBlockhash
}

func (b *Block) GetSize() int {
	return b.size
}

func (b *Block) GetTransactions() []Transaction {
	return b.transactions
}
