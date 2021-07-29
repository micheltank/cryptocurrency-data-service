package adapter

import (
	"context"
	"micheltank/cryptocurrency-data-service/internal/domain"
	so_chain "micheltank/cryptocurrency-data-service/internal/infra/so-chain"
	"sync"

	"github.com/pkg/errors"
)

type SoChainAdapterBlockchainApi struct {
	api so_chain.ISoChainBlockchainApi
}

func New(api so_chain.ISoChainBlockchainApi) domain.BlockchainApi {
	return &SoChainAdapterBlockchainApi{
		api: api,
	}
}

func (s *SoChainAdapterBlockchainApi) GetBlock(networkCode domain.NetworkCode, hash string) (domain.Block, error) {
	blockResponse, err := s.api.GetBlock(context.Background(), string(networkCode), hash)
	if err == so_chain.ErrNotFound {
		return domain.Block{}, ErrNotFound
	}
	if err != nil {
		return domain.Block{}, errors.Wrap(err, "failed to get block from SoChain api")
	}
	transactions, err := s.fetchFirstTenTransaction(networkCode, blockResponse.Data.Txs)
	if err != nil {
		return domain.Block{}, errors.Wrap(err, "failed to get transactions from SoChain api")
	}
	block := blockResponseToDomain(blockResponse, transactions)
	return block, nil
}

func (s *SoChainAdapterBlockchainApi) fetchFirstTenTransaction(networkCode domain.NetworkCode, txs []string) (domain.Transactions, error) {
	transactionsMap := make(map[string]domain.Transaction)
	wg := sync.WaitGroup{}
	wgDone := make(chan bool)
	cErrors := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxSize := 10
	if len(txs) < 10 {
		maxSize = len(txs)
	}
	for i := 0; i < maxSize; i++ {
		transactionId := txs[i]
		wg.Add(1)
		go func(id string) {
			defer wg.Done()

			transactionResponse, err := s.api.GetTransaction(ctx, string(networkCode), id)
			select {
			case <-ctx.Done():
				return
			default:
			}
			if err != nil {
				cancel()
				cErrors <- errors.Wrap(err, "failed to get transaction from SoChain api")
				return
			}
			transactionsMap[id], err = transactionResponseToDomain(transactionResponse)
			if err != nil {
				cancel()
				cErrors <- errors.Wrap(err, "failed to parse transaction from SoChain api")
			}
		}(transactionId)
	}
	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		var transactions domain.Transactions
		for _, transaction := range transactionsMap {
			transactions = append(transactions, transaction)
		}
		return transactions, nil
	case err := <-cErrors:
		close(cErrors)
		return nil, err
	}
}

func (s *SoChainAdapterBlockchainApi) GetTransaction(networkCode domain.NetworkCode, id string) (domain.Transaction, error) {
	transactionResponse, err := s.api.GetTransaction(context.Background(), string(networkCode), id)
	if err == so_chain.ErrNotFound {
		return domain.Transaction{}, ErrNotFound
	}
	if err != nil {
		return domain.Transaction{}, errors.Wrap(err, "failed to get transaction from SoChain api")
	}
	block, err := transactionResponseToDomain(transactionResponse)
	if err != nil {
		return domain.Transaction{}, errors.Wrap(err, "failed to parse transaction from SoChain api")
	}
	return block, nil
}
