package block

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

func MakeBlockHandler(routerGroup *gin.RouterGroup, service application.IService) {
	routerGroup.GET("/blocks/:networkCode/:hash", func(c *gin.Context) {
		V1GetBlock(c, service)
	})
}

// V1GetBlock godoc
// @Summary Get a block
// @Description Get a block along with the first ten transactions
// @ID get-block
// @Tags Blocks
// @Param networkCode path string true "The acronym of the network you're querying required"
// @Param hash path string true "The blockhash or height (number) on the network you're querying"
// @Produce json
// @Success 200 {object} presenter.BlockResponse
// @Error 400 {object} presenter.ApiError
// @Router /blocks/{networkCode}/{hash} [get]
func V1GetBlock(c *gin.Context, service application.IService) {
	networkCode := c.Param("networkCode")
	if networkCode == "" {
		c.Status(http.StatusBadRequest)
	}
	hash := c.Param("hash")
	if hash == "" {
		c.Status(http.StatusBadRequest)
	}
	block, err := service.GetBlock(domain.NetworkCode(networkCode), hash)
	if e, ok := err.(*domain.Error); ok && e != nil {
		logrus.WithError(err).Error("V1GetBlock returned client error")
		if errors.Cause(err) == adapter.ErrNotFound {
			c.JSON(http.StatusNotFound, presenter.NewApiError(*e))
			return
		}
		c.JSON(http.StatusBadRequest, presenter.NewApiError(*e))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("V1GetBlock returned InternalServerError")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, presenter.NewBlockResponse(block))
}