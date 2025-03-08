package handlers

import "backend-server/services"

type Handler struct {
	SmtpService services.ISmtpService
	JobService   services.IJobService
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SmtpHandler(smtp services.ISmtpService) *Handler {
	h.SmtpService = smtp
	return h
}

func (h *Handler) JobHandler(job services.IJobService) *Handler {
	h.JobService = job
	return h
}
