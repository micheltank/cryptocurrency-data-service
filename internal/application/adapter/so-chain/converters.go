package adapter

import (
	"micheltank/cryptocurrency-data-service/internal/domain"
	so_chain "micheltank/cryptocurrency-data-service/internal/infra/so-chain"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func blockResponseToDomain(block so_chain.BlockResponse, transactions domain.Transactions) domain.Block {
	return domain.NewBlock(domain.NetworkCode(block.Data.Network),
		block.Data.BlockNo,
		time.Unix(block.Data.Time, 0).UTC(),
		block.Data.PreviousBlockhash,
		block.Data.NextBlockhash,
		block.Data.Size,
		transactions)
}

func transactionResponseToDomain(transaction so_chain.TransactionResponse) (domain.Transaction, error) {
	sentValue, err := parseFloat(transaction.Data.SentValue)
	if err != nil {
		return domain.Transaction{}, err
	}
	fee, err := parseFloat(transaction.Data.Fee)
	if err != nil {
		return domain.Transaction{}, err
	}
	return domain.NewTransaction(transaction.Data.TxId,
		time.Unix(transaction.Data.Time, 0).UTC(),
		fee,
		sentValue), nil
}

func parseFloat(value string) (float64, error) {
	sentValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse transaction from SoChain api")
	}
	return sentValue, nil
}
