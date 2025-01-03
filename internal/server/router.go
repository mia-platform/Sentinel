package server

import (
	"context"
	"os"
	"path"
	"path/filepath"

	"github.com/mia-platform/sentinel/internal/config"
	"github.com/mia-platform/sentinel/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"

	glogrus "github.com/mia-platform/glogger/v4/loggers/logrus"
	middleware "github.com/mia-platform/glogger/v4/middleware/fiber"
	"github.com/sirupsen/logrus"
)


func NewApp(ctx context.Context, env config.EnvironmentVariables, log *logrus.Logger, cfg *config.Configuration) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	cmdName := filepath.Base(os.Args[0])
	middlewareLog := glogrus.GetLogger(logrus.NewEntry(log))
	app.Use(middleware.RequestMiddlewareLogger(middlewareLog, []string{"/-/"}))
	statusRoutes(app, cmdName, utils.ServiceVersionInformation())
	if env.ServicePrefix != "" && env.ServicePrefix != "/" {
		log.WithField("servicePrefix", env.ServicePrefix).Trace("applying service prefix")
		app.Use(pprof.New(pprof.Config{Prefix: path.Clean(env.ServicePrefix)}))
	}

	// Aggiungi qui le tue rotte

	return app, nil
}