package handler

import (
	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/service"
	"github.com/L1z1ng3r-sswe/instagram_clone/app/pkg/logging"
)

type Handler struct {
	service *service.Service
	log     *logging.Logger
}

func NewHandler(service *service.Service, log *logging.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}
