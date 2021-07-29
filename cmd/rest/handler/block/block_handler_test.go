package block_test

import (
	"encoding/json"
	blockHandler "micheltank/cryptocurrency-data-service/cmd/rest/handler/block"
	"micheltank/cryptocurrency-data-service/cmd/rest/presenter"
	adapter "micheltank/cryptocurrency-data-service/internal/application/adapter/so-chain"
	mock_application "micheltank/cryptocurrency-data-service/internal/application/mock"
	"micheltank/cryptocurrency-data-service/internal/domain"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
)

func TestV1GetBlock(t *testing.T) {

	t.Run("Regular get block", func(t *testing.T) {
		g := NewGomegaWithT(t)

		transactions := domain.Transactions{domain.NewTransaction("1", time.Now(), 0.05, 10)}
		block := domain.NewBlock("BTC", 1, time.Now(), "5", "7", 1, transactions)

		w, c := mockContext(string(block.GetNetworkCode()), "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf")
		service := mockService(t, block.GetNetworkCode(), block, nil)

		blockHandler.V1GetBlock(c, service)

		g.Expect(w.Code).Should(
			Equal(http.StatusOK))

		var got presenter.BlockResponse
		err := json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
		responseExpect := presenter.NewBlockResponse(block)
		g.Expect(responseExpect).Should(
			Equal(got))
	})

	t.Run("Get block with unsupported network code", func(t *testing.T) {
		g := NewGomegaWithT(t)

		transactions := domain.Transactions{domain.NewTransaction("1", time.Now(), 0.05, 10)}
		block := domain.NewBlock("ETH", 1, time.Now(), "5", "7", 1, transactions)

		w, c := mockContext(string(block.GetNetworkCode()), "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf")
		expectedError := block.GetNetworkCode().CheckIsSupported()
		service := mockService(t, block.GetNetworkCode(), domain.Block{}, expectedError)
		blockHandler.V1GetBlock(c, service)

		g.Expect(w.Code).Should(
			Equal(http.StatusBadRequest))

		var got presenter.ApiError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
		domainError := *expectedError.(*domain.Error)
		apiError := presenter.NewApiError(domainError)
		g.Expect(apiError).Should(
			Equal(got))
	})
	t.Run("Get block with absent hash", func(t *testing.T) {
		g := NewGomegaWithT(t)

		transactions := domain.Transactions{domain.NewTransaction("1", time.Now(), 0.05, 10)}
		block := domain.NewBlock("BTC", 1, time.Now(), "5", "7", 1, transactions)

		w, c := mockContext(string(block.GetNetworkCode()), "1234123")
		expectedError := domain.NewError(adapter.ErrNotFound, "block not found", "error.notFound", "")

		service := mockService(t, block.GetNetworkCode(), domain.Block{}, expectedError)
		blockHandler.V1GetBlock(c, service)

		g.Expect(w.Code).Should(
			Equal(http.StatusNotFound))

		var got presenter.ApiError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
		domainError := *expectedError.(*domain.Error)
		apiError := presenter.NewApiError(domainError)
		g.Expect(apiError).Should(
			Equal(got))
	})
}

func mockContext(networkCode, hash string) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{
		gin.Param{
			Key:   "networkCode",
			Value: networkCode,
		},
		gin.Param{
			Key:   "hash",
			Value: hash,
		}}
	return w, c
}

func mockService(t *testing.T, networkCode domain.NetworkCode, blockResponse domain.Block, errorResponse error) *mock_application.MockIService {
	ctrl := gomock.NewController(t)
	service := mock_application.NewMockIService(ctrl)
	service.EXPECT().
		GetBlock(gomock.Eq(networkCode), gomock.Any()).
		DoAndReturn(func(networkCode domain.NetworkCode, hash string) (domain.Block, error) {
			return blockResponse, errorResponse
		}).
		AnyTimes()
	return service
}
