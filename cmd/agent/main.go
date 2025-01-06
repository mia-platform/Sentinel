package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"

	"github.com/mia-platform/sentinel/internal/config"
	"github.com/mia-platform/sentinel/internal/monitor"
	"github.com/mia-platform/sentinel/internal/server"
)

func main() {
	// Carica le variabili d'ambiente
	envVars, err := env.ParseAs[config.EnvironmentVariables]()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	config, err := config.LoadServiceConfiguration(envVars.ConfigurationPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Gestione dei segnali
	sysChan := make(chan os.Signal, 1)
	signal.Notify(sysChan, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := monitor.New(*config)

	// Avvia monitoraggio
	go func() {
		if err := m.Start(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting monitor: %v\n", err)
			cancel()
		}
	}()

	exitCode := 0

	// Avvia il server REST
	if err = server.New(ctx, envVars, config, sysChan); err != nil {
		cancel()
		fmt.Fprintln(os.Stderr, err)
		exitCode = 1
	}

	<-sysChan
	fmt.Println("Sentinel agent terminated. Shutting down...")
	time.Sleep(2 * time.Second)

	close(sysChan)
	os.Exit(exitCode)
}
