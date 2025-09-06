package api

import (
	"messagebroker/internal/config"
	"messagebroker/internal/service"
)

type Handler struct {
	Config    config.Config
	Service   service.Servicer
}

type Option func(*Handler)

func WithConfig(config config.Config) Option {
	return func(s *Handler) {
		s.Config = config
	}
}

func WithSrvice(service service.Servicer) Option {
	return func(s *Handler) {
		s.Service = service
	}
}

func NewHandler(options ...Option) *Handler {
	h := &Handler{}
	for _, option := range options {
		option(h)
	}
	return h
}
