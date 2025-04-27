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

type ListFilesResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserAuthenticationResponse struct {
	Message string `json:"message"`
}

type RequestBodyResponse struct {
	Message string `json:"message"`
}

type ListRequestBodyResponse struct {
	Name       string `json:"name"`
	Promo_Text string `json:"promo_text"`
}

type UpdateRequestEntry struct {
	Name       string `json:"name"`
	Promo_Text string `json:"promo_text"`
}
