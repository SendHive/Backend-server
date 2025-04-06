package handlers

import "backend-server/services"

type Handler struct {
	SmtpService  services.ISmtpService
	JobService   services.IJobService
	UserService  services.IUser
	LoginService services.ILoginService
	FileService  services.IFileService
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

func (h *Handler) UserHandler(user services.IUser) *Handler {
	h.UserService = user
	return h
}

func (h *Handler) LoginHandler(login services.ILoginService) *Handler {
	h.LoginService = login
	return h
}

func (h *Handler) FileHeader(file services.IFileService) *Handler {
	h.FileService = file
	return h
}