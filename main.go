package main

import (
	"micheltank/cryptocurrency-data-service/cmd/rest"
	_ "micheltank/cryptocurrency-data-service/docs"
	"micheltank/cryptocurrency-data-service/internal/infra/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	appConfig := config.Env
	err := run(appConfig)
	if err != nil {
		logrus.WithError(err).
			Fatal("failed running application")
		return
	}
}

func run(appConfig config.Environment) error {
	logrus.Info("Starting application")

	// REST Server
	restApiServer, err := rest.NewServer(appConfig)
	if err != nil {
		return errors.Wrap(err, "failed to initialize restApiServer")
	}
	restApiErr := restApiServer.Run()
	logrus.Infof("Running http server on port %d", appConfig.Port)
	defer restApiServer.Shutdown()

	// Shutdown
	quit := notifyShutdown()
	select {
	case err := <-restApiErr:
		return errors.Wrap(err, "failed while running restApiServer")
	case <-quit:
		logrus.Info("Gracefully shutdown")
		return nil
	}
}

func notifyShutdown() chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	return quit
}