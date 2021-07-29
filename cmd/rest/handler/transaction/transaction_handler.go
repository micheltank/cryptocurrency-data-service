package transaction

import (
	"micheltank/cryptocurrency-data-service/cmd/rest/presenter"
	"micheltank/cryptocurrency-data-service/internal/application"
	adapter "micheltank/cryptocurrency-data-service/internal/application/adapter/so-chain"
	"micheltank/cryptocurrency-data-service/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func MakeTransactionHandler(routerGroup *gin.RouterGroup, service application.IService) {
	routerGroup.GET("/transactions/:networkCode/:transactionId", func(c *gin.Context) {
		V1GetTransaction(c, service)
	})
}

// V1GetTransaction godoc
// @Summary Get a transaction
// @Description Get a transaction
// @ID get-transaction
// @Tags Transactions
// @Param networkCode path string true "The acronym of the network you're querying required"
// @Param transactionId path string true "The transaction hash (id) on the network you're querying"
// @Produce json
// @Success 200 {object} presenter.TransactionResponse
// @Error 400 {object} presenter.ApiError
// @Router /transactions/{networkCode}/{transactionId} [get]
func V1GetTransaction(c *gin.Context, service application.IService) {
	networkCode := c.Param("networkCode")
	if networkCode == "" {
		c.Status(http.StatusBadRequest)
	}
	transactionId := c.Param("transactionId")
	if transactionId == "" {
		c.Status(http.StatusBadRequest)
	}
	transaction, err := service.GetTransaction(domain.NetworkCode(networkCode), transactionId)
	if e, ok := err.(*domain.Error); ok && e != nil {
		logrus.WithError(err).Error("V1GetTransaction returned client error")
		if errors.Cause(err) == adapter.ErrNotFound {
			c.JSON(http.StatusNotFound, presenter.NewApiError(*e))
			return
		}
		c.JSON(http.StatusBadRequest, presenter.NewApiError(*e))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("V1GetTransaction returned InternalServerError")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, presenter.NewTransactionResponse(transaction))
}
