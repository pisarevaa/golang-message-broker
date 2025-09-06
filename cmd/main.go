package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	docs "messagebroker/docs"

	"messagebroker/internal/config"
	"messagebroker/internal/api"
	"messagebroker/internal/router"
	"messagebroker/internal/service"
	"messagebroker/internal/logger"
)

const readTimeout = 5
const writeTimeout = 10
const shutdownTimeout = 10

// @title		Swagger Message Broker API
// @version	1.0
// @host		localhost:7000

func main() {
	ctxCancel, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctxStop, stop := signal.NotifyContext(ctxCancel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	logger.NewLogger()

	apiService := service.NewService(
		service.WithConfig(config),
	)

	docs.SwaggerInfo.Host = config.SwaggerURL

	handlers := api.NewHandler(
		api.WithConfig(config),
		api.WithSrvice(apiService),
	)
	srv := &http.Server{
		Addr:         config.Host,
		Handler:      router.NewRouter(handlers),
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
	}
	go func() {
		if errServer := srv.ListenAndServe(); errServer != nil && errServer != http.ErrServerClosed {
			log.Info().Msgf("could not listen on %s", config.Host)
		}
	}()
	log.Info().Msgf("api service is running on %s", config.Host)

	<-ctxStop.Done()
	shutdownCtx, timeout := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
	defer timeout()
	err = srv.Shutdown(shutdownCtx)
	if err != nil {
		log.Error().Err(err)
	}
	log.Info().Msg("api service is gracefully shutdown")
}
