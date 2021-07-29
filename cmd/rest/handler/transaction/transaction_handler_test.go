package transaction_test

import (
	"encoding/json"
	transactionHandler "micheltank/cryptocurrency-data-service/cmd/rest/handler/transaction"
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

func TestV1GetTransaction(t *testing.T) {

	t.Run("Regular get transaction", func(t *testing.T) {
		g := NewGomegaWithT(t)

		networkCode := domain.NetworkCode("BTC")
		transaction := domain.NewTransaction("1", time.Now(), 0.05, 10)

		w, c := mockContext(string(networkCode), "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf")
		service := mockService(t, networkCode, transaction, nil)

		transactionHandler.V1GetTransaction(c, service)

		g.Expect(w.Code).Should(
			Equal(http.StatusOK))

		var got presenter.TransactionResponse
		err := json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
		responseExpect := presenter.NewTransactionResponse(transaction)
		g.Expect(responseExpect).Should(
			Equal(got))
	})

	t.Run("Get transaction with unsupported network code", func(t *testing.T) {
		g := NewGomegaWithT(t)

		networkCode := domain.NetworkCode("ETH")

		w, c := mockContext(string(networkCode), "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf")
		expectedError := networkCode.CheckIsSupported()
		service := mockService(t, networkCode, domain.Transaction{}, expectedError)
		transactionHandler.V1GetTransaction(c, service)

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
	t.Run("Get transaction with absent network code", func(t *testing.T) {
		g := NewGomegaWithT(t)

		networkCode := domain.NetworkCode("BTC")

		w, c := mockContext(string(networkCode), "123213")
		expectedError := domain.NewError(adapter.ErrNotFound, "block not found", "error.notFound", "")
		service := mockService(t, networkCode, domain.Transaction{}, expectedError)
		transactionHandler.V1GetTransaction(c, service)

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

func mockContext(networkCode, transactionId string) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{
		gin.Param{
			Key:   "networkCode",
			Value: networkCode,
		},
		gin.Param{
			Key:   "transactionId",
			Value: transactionId,
		}}
	return w, c
}

func mockService(t *testing.T, networkCode domain.NetworkCode, transactionResponse domain.Transaction, errorResponse error) *mock_application.MockIService {
	ctrl := gomock.NewController(t)
	service := mock_application.NewMockIService(ctrl)
	service.EXPECT().
		GetTransaction(gomock.Eq(networkCode), gomock.Any()).
		DoAndReturn(func(networkCode domain.NetworkCode, transactionId string) (domain.Transaction, error) {
			return transactionResponse, errorResponse
		}).
		AnyTimes()
	return service
}
