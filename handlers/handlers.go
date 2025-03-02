package handlers

import "backend-server/services"

type Handler struct {
	SmtpService services.ISmtpService
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SmtpHandler(smtp services.ISmtpService) *Handler {
	h.SmtpService = smtp
	return h
}
