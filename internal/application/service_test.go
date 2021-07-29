package application_test

import (
	"micheltank/cryptocurrency-data-service/internal/application"
	"micheltank/cryptocurrency-data-service/internal/domain"
	mock_domain "micheltank/cryptocurrency-data-service/internal/domain/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
)

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	blockchainApi := mock_domain.NewMockBlockchainApi(ctrl)

	t.Run("Regular get block", func(t *testing.T) {
		g := NewGomegaWithT(t)

		networkCode := domain.NetworkCode("BTC")
		transactions := domain.Transactions{domain.NewTransaction("1", time.Now(), 0.05, 10)}
		expectedBlock := domain.NewBlock(networkCode, 1, time.Now(), "5", "7", 1, transactions)

		blockchainApi.EXPECT().
			GetBlock(gomock.Any(), gomock.Any()).
			DoAndReturn(func(networkCode domain.NetworkCode, hash string) (domain.Block, error) {
				return expectedBlock, nil
			}).
			AnyTimes()
		service := application.NewService(blockchainApi)

		block, err := service.GetBlock(networkCode, "21312")
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
		g.Expect(block.GetDateTime()).Should(
			Equal(expectedBlock.GetDateTime()))
		g.Expect(block.GetTransactions()).Should(
			Equal(expectedBlock.GetTransactions()))
	})
	t.Run("Regular get transaction", func(t *testing.T) {
		g := NewGomegaWithT(t)

		networkCode := domain.NetworkCode("BTC")
		expectedTransaction := domain.NewTransaction("1", time.Now(), 0.05, 10)

		blockchainApi.EXPECT().
			GetTransaction(gomock.Any(), gomock.Any()).
			DoAndReturn(func(networkCode domain.NetworkCode, transactionId string) (domain.Transaction, error) {
				return expectedTransaction, nil
			}).
			AnyTimes()
		service := application.NewService(blockchainApi)

		transaction, err := service.GetTransaction(networkCode, "21312")
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(transaction).Should(
			Not(BeNil()))
		g.Expect(transaction.GetTransactionId()).Should(
			Equal(expectedTransaction.GetTransactionId()))
		g.Expect(transaction.GetFee()).Should(
			Equal(expectedTransaction.GetFee()))
		g.Expect(transaction.GetDateTime()).Should(
			Equal(expectedTransaction.GetDateTime()))
		g.Expect(transaction.GetSentValue()).Should(
			Equal(expectedTransaction.GetSentValue()))
	})
}
