package application

import (
	adapter "micheltank/cryptocurrency-data-service/internal/application/adapter/so-chain"
	"micheltank/cryptocurrency-data-service/internal/domain"
)

type IService interface {
	GetBlock(networkCode domain.NetworkCode, hash string) (domain.Block, error)
	GetTransaction(networkCode domain.NetworkCode, id string) (domain.Transaction, error)
}

type Service struct {
	blockchainApi domain.BlockchainApi
}

func NewService(blockchainApi domain.BlockchainApi) IService {
	return &Service{
		blockchainApi: blockchainApi,
	}
}

func (s *Service) GetBlock(networkCode domain.NetworkCode, hash string) (domain.Block, error) {
	err := networkCode.CheckIsSupported()
	if err != nil {
		return domain.Block{}, err
	}
	block, err := s.blockchainApi.GetBlock(networkCode, hash)
	if err == adapter.ErrNotFound {
		return domain.Block{}, domain.NewError(err, "block not found", "error.notFound", "")
	}
	if err != nil {
		return domain.Block{}, err
	}
	return block, nil
}

func (s *Service) GetTransaction(networkCode domain.NetworkCode, id string) (domain.Transaction, error) {
	err := networkCode.CheckIsSupported()
	if err != nil {
		return domain.Transaction{}, err
	}
	block, err := s.blockchainApi.GetTransaction(networkCode, id)
	if err == adapter.ErrNotFound {
		return domain.Transaction{}, domain.NewError(err, "block not found", "error.notFound", "")
	}
	if err != nil {
		return domain.Transaction{}, err
	}
	return block, nil
}
