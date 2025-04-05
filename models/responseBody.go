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

type CreateJobResponse struct {
	Message string `json:"message"`
	TaskId  string `json:"task_id"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
}

type CreateLoginResponse struct {
	Message string `json:"message"`
}

type ListJobEntryResponse struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

type CreateFileEntryResponse struct {
	Message string `json:"message"`
}