package adapter_test

import (
	"context"
	"fmt"
	adapter "micheltank/cryptocurrency-data-service/internal/application/adapter/so-chain"
	"micheltank/cryptocurrency-data-service/internal/domain"
	so_chain "micheltank/cryptocurrency-data-service/internal/infra/so-chain"
	mock_so_chain "micheltank/cryptocurrency-data-service/internal/infra/so-chain/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
)

func TestAdapter(t *testing.T) {
	ctrl := gomock.NewController(t)
	blockchainApi := mock_so_chain.NewMockISoChainBlockchainApi(ctrl)

	t.Run("Regular get block", func(t *testing.T) {
		g := NewGomegaWithT(t)

		networkCode := domain.NetworkCode("BTC")
		timeNow := time.Now()
		transaction := domain.NewTransaction("1", timeNow, 0.05, 10)
		transactions := domain.Transactions{transaction}
		expectedBlock := domain.NewBlock(networkCode, 1, timeNow, "5", "7", 1, transactions)

		blockResponse := so_chain.BlockResponse{
			Status: "success",
			Data:   so_chain.BlockDataResponse{
				Network: string(expectedBlock.GetNetworkCode()),
				Blockhash:         "123213",
				BlockNo:           expectedBlock.GetBlockNumber(),
				MiningDifficulty:  "123123",
				Time:              timeNow.Unix(),
				Confirmations:     0,
				IsOrphan:          false,
				Txs:               []string{transactions[0].GetTransactionId()},
				Merkleroot:        "",
				PreviousBlockhash: expectedBlock.GetPreviousBlockhash(),
				NextBlockhash:     expectedBlock.GetNextBlockhash(),
				Size:              expectedBlock.GetSize(),
			},
		}
		transactionResponse := so_chain.TransactionResponse{
			Status: "success",
			Data:   so_chain.TransactionDataResponse{
				TxId:      transaction.GetTransactionId(),
				Time:      timeNow.Unix(),
				SentValue: fmt.Sprintf("%f", transaction.GetSentValue()),
				Fee:       fmt.Sprintf("%f", transaction.GetFee()),
			},
		}
		blockchainApi.EXPECT().
			GetBlock(gomock.Any(), gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, networkCode string, hash string) (so_chain.BlockResponse, error) {
				return blockResponse, nil
			}).
			AnyTimes()
		blockchainApi.EXPECT().
			GetTransaction(gomock.Any(), gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, networkCode string, transactionId string) (so_chain.TransactionResponse, error) {
				return transactionResponse, nil
			}).
			AnyTimes()
		adapter := adapter.New(blockchainApi)

		block, err := adapter.GetBlock(networkCode, "21312")
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(block).Should(
			Not(BeNil()))
		g.Expect(block.GetNetworkCode()).Should(
			Equal(expectedBlock.GetNetworkCode()))
		g.Expect(block.GetBlockNumber()).Should(
			Equal(expectedBlock.GetBlockNumber()))
		g.Expect(block.GetSize()).Should(
			Equal(expectedBlock.GetSize()))
		g.Expect(block.GetNextBlockhash()).Should(
			Equal(expectedBlock.GetNextBlockhash()))
		g.Expect(block.GetPreviousBlockhash()).Should(
			Equal(expectedBlock.GetPreviousBlockhash()))
		g.Expect(block.GetDateTime().Unix()).Should(
			Equal(block.GetDateTime().Unix()))
		for i, transaction := range block.GetTransactions() {
			g.Expect(transaction.GetTransactionId()).Should(
				Equal(expectedBlock.GetTransactions()[i].GetTransactionId()))
			g.Expect(transaction.GetFee()).Should(
				Equal(expectedBlock.GetTransactions()[i].GetFee()))
			g.Expect(transaction.GetSentValue()).Should(
				Equal(expectedBlock.GetTransactions()[i].GetSentValue()))
			g.Expect(transaction.GetDateTime().Unix()).Should(
				Equal(expectedBlock.GetTransactions()[i].GetDateTime().Unix()))
		}
	})
	t.Run("Regular get transaction", func(t *testing.T) {
		g := NewGomegaWithT(t)

		networkCode := domain.NetworkCode("BTC")
		expectedTransaction := domain.NewTransaction("1", time.Now(), 0.05, 10)

		timeNow := time.Now()
		transaction := domain.NewTransaction("1", timeNow, 0.05, 10)
		transactionResponse := so_chain.TransactionResponse{
			Status: "success",
			Data:   so_chain.TransactionDataResponse{
				TxId:      transaction.GetTransactionId(),
				Time:      timeNow.Unix(),
				SentValue: fmt.Sprintf("%f", transaction.GetSentValue()),
				Fee:       fmt.Sprintf("%f", transaction.GetFee()),
			},
		}
		blockchainApi.EXPECT().
			GetTransaction(gomock.Any(), gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, networkCode string, transactionId string) (so_chain.TransactionResponse, error) {
				return transactionResponse, nil
			}).
			AnyTimes()
		adapter := adapter.New(blockchainApi)

		transaction, err := adapter.GetTransaction(networkCode, "21312")
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(transaction).Should(
			Not(BeNil()))
		g.Expect(transaction.GetTransactionId()).Should(
			Equal(expectedTransaction.GetTransactionId()))
		g.Expect(transaction.GetFee()).Should(
			Equal(expectedTransaction.GetFee()))
		g.Expect(transaction.GetDateTime().Unix()).Should(
			Equal(expectedTransaction.GetDateTime().Unix()))
		g.Expect(transaction.GetSentValue()).Should(
			Equal(expectedTransaction.GetSentValue()))
	})
}

