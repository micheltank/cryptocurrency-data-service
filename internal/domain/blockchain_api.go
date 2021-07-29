package domain

type BlockchainApi interface {
	GetBlock(networkCode NetworkCode, hash string) (Block, error)
	GetTransaction(networkCode NetworkCode, id string) (Transaction, error)
}
