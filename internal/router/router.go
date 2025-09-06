package router

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"messagebroker/internal/api"
)

// Создание роутера.
func NewRouter(handlers *api.Handler) http.Handler{
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/queues/{queue_name}/messages", handlers.SendMessage)
	mux.HandleFunc("GET /v1/queues/{queue_name}/messages", handlers.GetMessage)
	mux.HandleFunc("POST /v1/queues/{queue_name}/subscriptions", handlers.Subscribe)
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	var handler http.Handler = mux
	handler = api.RecoveryMiddleware(handler)
	handler = api.NewLoggingMiddleware()(handler)

	return handler
}