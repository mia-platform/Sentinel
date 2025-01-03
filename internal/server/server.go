package server

import (
	"context"
	"fmt"

	"github.com/mia-platform/sentinel/internal/config"

	glogrus "github.com/mia-platform/glogger/v4/loggers/logrus"
)

func New[Signal any](ctx context.Context, envVars config.EnvironmentVariables, cfg *config.Configuration, sysChannel <-chan Signal) error {
	// Init logger instance.
	log, err := glogrus.InitHelper(glogrus.InitOptions{Level: envVars.LogLevel})
	if err != nil {
		panic(err)
	}

	app, err := NewApp(ctx, envVars, log, cfg)
	if err != nil {
		return fmt.Errorf("error creating app: %w", err)
	}

	go func() {
		log.WithField("port", envVars.HTTPPort).Info("starting server")
		if err := app.Listen(fmt.Sprintf("%s:%s", envVars.HTTPAddress, envVars.HTTPPort)); err != nil {
			log.Println(err)
		}
	}()

	<-sysChannel

	if err := app.Shutdown(); err != nil {
		return err
	}

	return nil
}
