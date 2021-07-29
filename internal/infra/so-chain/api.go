package so_chain

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type ISoChainBlockchainApi interface {
	GetBlock(ctx context.Context, networkCode string, hash string) (BlockResponse, error)
	GetTransaction(ctx context.Context, networkCode string, id string) (TransactionResponse, error)
}

type SoChainBlockchainApi struct {
	host   string
	client *http.Client
}

func New(host string) ISoChainBlockchainApi {
	return &SoChainBlockchainApi{
		host: host,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s SoChainBlockchainApi) GetBlock(ctx context.Context, networkCode string, hash string) (BlockResponse, error) {
	url := fmt.Sprintf("%s/api/v2/get_block/%s/%s", s.host, networkCode, hash)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return BlockResponse{}, errors.Wrap(err, "failed to get block from SoChain api")
	}
	req = req.WithContext(ctx)
	resp, err := s.client.Do(req)
	if err != nil {
		return BlockResponse{}, errors.Wrap(err, "failed to get block from SoChain api")
	}
	if resp.StatusCode == http.StatusNotFound {
		return BlockResponse{}, ErrNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return BlockResponse{}, errors.New(fmt.Sprintf("SoChaine api returned %d status code error", resp.StatusCode))
	}
	var block BlockResponse
	err = json.NewDecoder(resp.Body).Decode(&block)
	if err != nil {
		return BlockResponse{}, errors.Wrap(err, "failed to decode block from SoChain api")
	}
	return block, nil
}

func (s SoChainBlockchainApi) GetTransaction(ctx context.Context, networkCode string, transactionId string) (TransactionResponse, error) {
	url := fmt.Sprintf("%s/api/v2/tx/%s/%s", s.host, networkCode, transactionId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TransactionResponse{}, errors.Wrap(err, "failed to get transaction from SoChain api")
	}
	req = req.WithContext(ctx)
	resp, err := s.client.Do(req)
	if err != nil {
		return TransactionResponse{}, errors.Wrap(err, "failed to get transaction from SoChain api")
	}
	if resp.StatusCode == http.StatusNotFound {
		return TransactionResponse{}, ErrNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return TransactionResponse{}, errors.New(fmt.Sprintf("SoChaine api returned %d status code error", resp.StatusCode))
	}
	var transaction TransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&transaction)
	if err != nil {
		return TransactionResponse{}, errors.Wrap(err, "failed to decode transaction from SoChain api")
	}
	return transaction, nil
}
