package service

import (
	"messagebroker/internal/config"
)

type Service struct {
	Config   config.Config
}

type Option func(*Service)

func WithConfig(config config.Config) Option {
	return func(s *Service) {
		s.Config = config
	}
}

func NewService(options ...Option) *Service {
	h := &Service{}
	for _, option := range options {
		option(h)
	}
	return h
}
