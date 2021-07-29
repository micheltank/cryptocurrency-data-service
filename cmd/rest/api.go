package rest

import (
	"context"
	"fmt"
	"micheltank/cryptocurrency-data-service/cmd/rest/handler"
	"micheltank/cryptocurrency-data-service/cmd/rest/handler/block"
	"micheltank/cryptocurrency-data-service/cmd/rest/handler/transaction"
	"micheltank/cryptocurrency-data-service/internal/application"
	so_chain_adapter "micheltank/cryptocurrency-data-service/internal/application/adapter/so-chain"
	"micheltank/cryptocurrency-data-service/internal/infra/config"
	so_chain "micheltank/cryptocurrency-data-service/internal/infra/so-chain"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Api struct {
	httpServer *http.Server
}

// NewServer godoc
// @title Blockchain API
// @BasePath /api/v1
// @version 1.0
func NewServer(config config.Environment) (*Api, error) {
	router := gin.Default()
	base := router.Group("/api")

	v1 := base.Group("/v1")

	// di
	soChainApi := so_chain.New(config.SoChainApiHost)
	blockchainApi := so_chain_adapter.New(soChainApi)
	service := application.NewService(blockchainApi)

	// handlers
	handler.MakeHealthCheckHandler(base)
	block.MakeBlockHandler(v1, service)
	transaction.MakeTransactionHandler(v1, service)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	httpServer := &http.Server{Addr: fmt.Sprintf(":%d", config.Port), Handler: router}

	return &Api{
		httpServer: httpServer,
	}, nil
}

func (api *Api) Run() <-chan error {
	out := make(chan error)
	go func() {
		if err := api.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			out <- errors.Wrap(err, "failed to listen and serve api")
		}
	}()
	return out
}

func (api *Api) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := api.httpServer.Shutdown(ctx)
		if err != nil {
			logrus.
				WithError(err).
				Error("Server forced to shutdown")
		}
	}()
}
