package models

type CreateSmtpEntryResponse struct {
	Message string `json:"message"`
}

type UpdateSmtpEntryResponse struct {
	Message string `json:"message"`
}

type ListSmtpEntryResponse struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Username string `json:"username"`
}
